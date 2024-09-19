package actuatorfx

import (
	"github.com/caarlos0/env/v11"
	"go.uber.org/fx"

	"github.com/hadroncorp/geck/actuator"
)

func AsActuator(t any) any {
	return fx.Annotate(t,
		fx.As(new(actuator.Actuator)),
		fx.ResultTags(`group:"actuators"`),
	)
}

var ActuatorModule = fx.Module("actuator",
	fx.Provide(
		env.ParseAs[actuator.ConfigManager],
		env.ParseAs[actuator.ConfigDiskActuator],
		AsActuator(actuator.NewDiskActuator),
		actuator.NewManager,
	),
)
