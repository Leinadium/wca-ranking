package fx

import (
	uberfx "go.uber.org/fx"
	"leinadium.dev/wca-ranking/internal/core/service"
)

var (
	ServicesModule = uberfx.Module("service",
		uberfx.Provide(service.NewFileService),
		uberfx.Provide(service.NewPersonService),
		uberfx.Provide(service.NewRankingService),
		uberfx.Provide(service.NewSearchService),
		uberfx.Provide(service.NewStateService),
		uberfx.Provide(service.NewUserService),
	)
)
