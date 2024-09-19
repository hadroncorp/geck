package identifierfx

import (
	"go.uber.org/fx"

	"github.com/hadroncorp/geck/identifier"
)

var IdentifierKSUIDModule = fx.Module("identifier_ksuid",
	fx.Provide(
		fx.Annotate(
			identifier.NewFactoryKSUID,
			fx.As(new(identifier.Factory)),
		),
	),
)

var IdentifierUUIDModule = fx.Module("identifier_uuid",
	fx.Provide(
		fx.Annotate(
			identifier.NewFactoryUUID,
			fx.As(new(identifier.Factory)),
		),
	),
)
