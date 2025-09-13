package fx

import (
	"go.uber.org/fx"

	"leinadium.dev/wca-ranking/internal/adapter/wca"
	"leinadium.dev/wca-ranking/internal/core/port"
)

var WCAModule = fx.Module("wca",
	fx.Provide(wca.NewRequester),
	fx.Provide(fx.Annotate(wca.NewWCAAPIRequester, fx.As(new(port.WCAAPIRequester)))),
	// fx.Provide(fx.Annotate(wca.NewWCAAPIService, fx.As(new(port.WCAAPIService)))),
	// fx.Provide(fx.Annotate(wca.NewWCATokenService, fx.As(new(port.WCATokenService)))),
	// fx.Provide(fx.Annotate(wca.NewWCATokenRequester, fx.As(new(port.WCATokenRequester)))),
)
