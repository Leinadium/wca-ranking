package fx

import (
	"go.uber.org/fx"

	"leinadium.dev/wca-ranking/internal/adapter/storage/mysql"
	"leinadium.dev/wca-ranking/internal/adapter/storage/mysql/repository"
	"leinadium.dev/wca-ranking/internal/adapter/storage/mysql/schema"
)

var (
	StorageModule = fx.Module("storage",
		fx.Provide(repository.NewPersonRepository),
		fx.Provide(repository.NewRankingRepository),
		fx.Provide(repository.NewSearchRepository),
		fx.Provide(repository.NewStateRepository),
		fx.Provide(repository.NewUserRepository),

		fx.Provide(mysql.New),
		fx.Provide(func(db *mysql.DB) *schema.Queries { return schema.New(db.DB) }),
	)
)
