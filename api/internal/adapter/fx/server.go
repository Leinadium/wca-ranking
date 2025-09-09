package fx

import (
	"go.uber.org/fx"
	"leinadium.dev/wca-ranking/internal/adapter/config"
	"leinadium.dev/wca-ranking/internal/adapter/server"
	"leinadium.dev/wca-ranking/internal/adapter/server/handler"
	"leinadium.dev/wca-ranking/internal/adapter/server/handler/routes"
)

var (
	ServerModule = fx.Module("server",
		fx.Provide(NewFXServer),
		fx.Provide(routes.NewStatesGroup),

		// start
		fx.Invoke(func(*server.Server) {}),
	)
)

func NewFXServer(lc fx.Lifecycle, config *config.Server, handlers []handler.Handler, groups []handler.HandlerGroup) *server.Server {
	srv := server.NewServer(config, handlers, groups)
	lc.Append(fx.Hook{OnStart: srv.Run, OnStop: srv.Stop})
	return srv
}
