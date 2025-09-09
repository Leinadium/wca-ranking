package fx

import (
	"go.uber.org/fx"
	"leinadium.dev/wca-ranking/internal/adapter/server"
	"leinadium.dev/wca-ranking/internal/adapter/server/handler/routes"
)

var (
	ServerModule = fx.Module("server",
		fx.Provide(server.NewServer),
		fx.Provide(routes.NewStatesGroup),
	)
)

func NewFXServer(lc fx.Lifecycle, sv *server.Server) {
	lc.Append(fx.Hook{
		OnStart: func (ctx func(context.Context) error {
			go sv.Run()
		}),
	})
}