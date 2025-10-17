package fx

import (
	"go.uber.org/fx"
	"leinadium.dev/wca-ranking/internal/adapter/config"
)

var ConfigModule = fx.Module("config",
	fx.Provide(config.New),
	fx.Provide(func(c *config.Config) *config.Server { return &c.Server }),
	fx.Provide(func(c *config.Config) *config.DB { return &c.DB }),
	fx.Provide(func(c *config.Config) *config.WCA { return &c.WCA }),
	fx.Provide(func(c *config.Config) *config.Auth { return &c.Auth }),
	fx.Provide(func(c *config.Config) *config.Updater { return &c.Updater }),
	fx.Supply(config.ConfigFile("./config.toml")),
)
