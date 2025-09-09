package fx

import (
	"go.uber.org/fx"

	"leinadium.dev/wca-ranking/internal/adapter/wca"
)

var (
	WCAModule = fx.Module("wca",
		fx.Provide(wca.NewRequester),
		fx.Provide(wca.NewWCAAPIRequester),
		fx.Provide(wca.NewWCATokenService),
	)
)
