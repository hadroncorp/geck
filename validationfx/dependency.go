package validationfx

import (
	"go.uber.org/fx"

	"github.com/hadroncorp/geck/validation"
)

var GoPlaygroundValidationModule = fx.Module("validation_go_playground",
	fx.Provide(
		fx.Annotate(
			validation.NewGoPlaygroundValidator,
			fx.As(new(validation.Validator)),
		),
	),
)
