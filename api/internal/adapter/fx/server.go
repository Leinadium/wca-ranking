package fx

import (
	"go.uber.org/fx"
	"leinadium.dev/wca-ranking/internal/adapter/config"
	"leinadium.dev/wca-ranking/internal/adapter/server"
	"leinadium.dev/wca-ranking/internal/adapter/server/handler"
)

type ServerParams struct {
	fx.In

	Config  *config.Server
	Handler *handler.ServerHandler
}

func NewFXServer(lc fx.Lifecycle, p ServerParams) *server.Server {
	srv := server.NewServer(p.Config, p.Handler)
	lc.Append(fx.Hook{OnStart: srv.Run, OnStop: srv.Stop})
	return srv
}

var ServerModule = fx.Module("server2",
	fx.Provide(NewFXServer),
	fx.Provide(handler.NewServerHandler),
	fx.Invoke(func(*server.Server) {}),
)
