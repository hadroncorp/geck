package main

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"

	"github.com/hadroncorp/geck/actuatorfx"
	"github.com/hadroncorp/geck/applicationfx"
	gecksql "github.com/hadroncorp/geck/data/sql"
	"github.com/hadroncorp/geck/datafx/sqlfx"
	"github.com/hadroncorp/geck/observability/loggingfx"
	"github.com/hadroncorp/geck/securityfx"
	"github.com/hadroncorp/geck/transportfx"
	"github.com/hadroncorp/geck/validationfx"
	"http-server-db/taskfx"
)

func NewPostgres(lifecycle fx.Lifecycle, cfg gecksql.Config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.ConnectionString)
	if err != nil {
		return nil, err
	}
	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})
	return db, nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Warn().Err(err).Msg("error loading .env file, using system env variables")
	}

	app := fx.New(
		fx.NopLogger, // comment this if you want detailed dep injection info from Uber FX
		applicationfx.ApplicationModule,
		loggingfx.ZerologLoggerModule,
		actuatorfx.ActuatorModule,
		validationfx.GoPlaygroundValidationModule,
		securityfx.CognitoModule,
		transportfx.TransportModuleHTTP,
		transportfx.TransportJWTModuleHTTP,
		fx.Provide(fx.Annotate(
			NewPostgres,
			fx.As(new(gecksql.Client)),
		)),
		sqlfx.MainModule,
		sqlfx.TransactionModule,
		sqlfx.DefaultDecorators,
		taskfx.TaskModule,
	)
	app.Run()
}
