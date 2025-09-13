package fx

import (
	"go.uber.org/fx"

	"leinadium.dev/wca-ranking/internal/adapter/storage/mysql"
	"leinadium.dev/wca-ranking/internal/adapter/storage/mysql/repository"
	"leinadium.dev/wca-ranking/internal/adapter/storage/mysql/schema"
	"leinadium.dev/wca-ranking/internal/core/port"
)

var StorageModule = fx.Module("storage",
	fx.Provide(fx.Annotate(repository.NewPersonRepository, fx.As(new(port.PersonRepository)))),
	fx.Provide(fx.Annotate(repository.NewRankingRepository, fx.As(new(port.RankingRepository)))),
	fx.Provide(fx.Annotate(repository.NewSearchRepository, fx.As(new(port.SearchRepository)))),
	fx.Provide(fx.Annotate(repository.NewStateRepository, fx.As(new(port.StateRepository)))),
	fx.Provide(fx.Annotate(repository.NewUserRepository, fx.As(new(port.UserRepository)))),

	fx.Provide(mysql.New),
	fx.Provide(mysql.Schema),
	fx.Provide(schema.New),
)
