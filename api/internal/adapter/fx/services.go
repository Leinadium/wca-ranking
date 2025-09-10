package fx

import (
	"go.uber.org/fx"
	"leinadium.dev/wca-ranking/internal/core/port"
	"leinadium.dev/wca-ranking/internal/core/service"
)

var (
	ServicesModule = fx.Module("service",
		fx.Provide(fx.Annotate(service.NewFileService, fx.As(new(port.FileService)))),
		fx.Provide(fx.Annotate(service.NewPersonService, fx.As(new(port.PersonService)))),
		fx.Provide(fx.Annotate(service.NewRankingService, fx.As(new(port.RankingService)))),
		fx.Provide(fx.Annotate(service.NewSearchService, fx.As(new(port.SearchService)))),
		fx.Provide(fx.Annotate(service.NewStateService, fx.As(new(port.StateService)))),
		fx.Provide(fx.Annotate(service.NewUserService, fx.As(new(port.UserService)))),
	)
)
