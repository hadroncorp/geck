package sqlfx

import (
	"github.com/caarlos0/env/v11"
	"go.uber.org/fx"

	"github.com/hadroncorp/geck/actuatorfx"
	"github.com/hadroncorp/geck/data/persistence"
	gecksql "github.com/hadroncorp/geck/data/sql"
	"github.com/hadroncorp/geck/observability/logging"
)

var MainModule = fx.Module("sql",
	fx.Provide(
		env.ParseAs[gecksql.Config],
		actuatorfx.AsActuator(gecksql.NewActuator),
	),
)

var TransactionModule = fx.Module("sql",
	fx.Provide(
		env.ParseAs[gecksql.ConfigTransactionFactory],
		fx.Annotate(
			gecksql.NewTransactionContextFactory,
			fx.As(new(persistence.TransactionContextFactory)),
		),
	),
)

var DefaultDecorators = fx.Decorate(
	func(src gecksql.Client, logger logging.Logger, factory persistence.TransactionContextFactory) gecksql.Client {
		src = TransactionalDecorator(src, logger, factory)
		return LoggerDecorator(src, logger)
	},
)

var LoggerDecorator = func(src gecksql.Client, logger logging.Logger) gecksql.Client {
	return gecksql.NewLoggerClient(logger, src)
}

var TransactionalDecorator = func(src gecksql.Client, logger logging.Logger, factory persistence.TransactionContextFactory) gecksql.Client {
	factorySQL, ok := factory.(gecksql.TransactionContextFactory)
	if !ok {
		return src
	}
	return gecksql.NewTransactionalClient(factorySQL, logger, src)
}
