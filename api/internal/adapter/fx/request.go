package fx

import (
	"go.uber.org/fx"

	"leinadium.dev/wca-ranking/internal/adapter/request"
	"leinadium.dev/wca-ranking/internal/adapter/request/auth"
	"leinadium.dev/wca-ranking/internal/adapter/request/user"
	"leinadium.dev/wca-ranking/internal/adapter/request/wca"
	"leinadium.dev/wca-ranking/internal/core/port"
)

var WCAModule = fx.Module("request",
	fx.Provide(request.NewRequester),
	fx.Provide(fx.Annotate(wca.NewWCAAPIRequester, fx.As(new(port.WCAAPIRequester)))),
	fx.Provide(fx.Annotate(user.NewUserRequester, fx.As(new(port.UserRequester)))),
	fx.Provide(fx.Annotate(auth.NewAuthRequester, fx.As(new(port.AuthRequester)))),
)
