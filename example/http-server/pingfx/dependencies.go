package pingfx

import (
	"go.uber.org/fx"

	"github.com/hadroncorp/geck/observability/loggingfx"
	"github.com/hadroncorp/geck/transportfx"
	"http-server/ping"
)

var PingModule = fx.Module("ping",
	fx.Decorate(
		loggingfx.DecorateLoggerWithModule("ping"),
	),
	fx.Provide(
		transportfx.AsVersionedControllerHTTP(ping.NewControllerHTTP),
	),
)
