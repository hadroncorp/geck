package taskfx

import (
	"go.uber.org/fx"

	"github.com/hadroncorp/geck/transportfx"
	"http-server-db/task"
)

var TaskModule = fx.Module("task",
	fx.Provide(
		fx.Annotate(
			task.NewRepositorySQL,
			fx.As(new(task.Repository)),
		),
		task.NewService,
		transportfx.AsVersionedControllerHTTP(task.NewControllerHTTP),
	),
)
