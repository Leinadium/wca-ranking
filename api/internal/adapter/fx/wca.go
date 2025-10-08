package fx

import (
	"go.uber.org/fx"

	"leinadium.dev/wca-ranking/internal/adapter/request"
	"leinadium.dev/wca-ranking/internal/adapter/request/wca"
	"leinadium.dev/wca-ranking/internal/core/port"
)

var WCAModule = fx.Module("wca",
	fx.Provide(request.NewRequester),
	fx.Provide(fx.Annotate(wca.NewWCAAPIRequester, fx.As(new(port.WCAAPIRequester)))),
	// fx.Provide(fx.Annotate(wca.NewWCAAPIService, fx.As(new(port.WCAAPIService)))),
	// fx.Provide(fx.Annotate(wca.NewWCATokenService, fx.As(new(port.WCATokenService)))),
	// fx.Provide(fx.Annotate(wca.NewWCATokenRequester, fx.As(new(port.WCATokenRequester)))),
)
