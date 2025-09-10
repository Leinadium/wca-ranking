package fx

import (
	"go.uber.org/fx"
	"leinadium.dev/wca-ranking/internal/adapter/config"
	"leinadium.dev/wca-ranking/internal/adapter/server"
	"leinadium.dev/wca-ranking/internal/adapter/server/handler"
	"leinadium.dev/wca-ranking/internal/adapter/server/handler/routes"
)

type ServerParams struct {
	fx.In

	Config   *config.Server
	Handlers []handler.Handler       `group:"handlers"`
	Groups   []*handler.HandlerGroup `group:"groups"`
}

func NewFXServer(lc fx.Lifecycle, p ServerParams) *server.Server {
	srv := server.NewServer(p.Config, p.Handlers, p.Groups)
	lc.Append(fx.Hook{OnStart: srv.Run, OnStop: srv.Stop})
	return srv
}

func AsGroup(f any) any {
	return fx.Annotate(f, fx.ResultTags(`group:"groups"`))
}

func AsHandler(f any) any {
	return fx.Annotate(f, fx.As(new(handler.Handler)), fx.ResultTags(`group:"handlers"`))
}

var ServerModule = fx.Module("server",
	fx.Provide(NewFXServer),
	fx.Provide(AsGroup(routes.NewStatesGroup)),

	// start
	fx.Invoke(func(*server.Server) {}),
)
