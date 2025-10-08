package fx

import (
	"go.uber.org/fx"
	"leinadium.dev/wca-ranking/internal/adapter/config"
	"leinadium.dev/wca-ranking/internal/adapter/updater"
	"leinadium.dev/wca-ranking/internal/core/port"
)

type UpdaterParams struct {
	fx.In

	Remote port.RemoteService
	Sync   port.SyncService
	File   port.FileService
	Config *config.Updater
}

func NewFXUpdater(lc fx.Lifecycle, p UpdaterParams) *updater.Updater {
	up := updater.NewUpdater(
		p.Remote,
		p.Sync,
		p.File,
		p.Config,
	)
	lc.Append(fx.Hook{OnStart: up.Start, OnStop: up.Stop})
	return up
}

var UpdateModule = fx.Module("update",
	fx.Provide(NewFXUpdater),
	fx.Invoke(func(*updater.Updater) {}),
)
