package updater

import (
	"context"
	"log"
	"strings"
	"time"

	"leinadium.dev/wca-ranking/internal/adapter/config"
	"leinadium.dev/wca-ranking/internal/core/port"
)

type Updater struct {
	remote port.RemoteService
	sync   port.SyncService
	file   port.FileService
	config *config.Updater

	timer *time.Ticker
	stop  chan struct{}
}

func NewUpdater(
	remote port.RemoteService,
	sync port.SyncService,
	file port.FileService,
	config *config.Updater,
) *Updater {
	return &Updater{
		remote: remote,
		sync:   sync,
		file:   file,
		config: config,

		timer: nil,
		stop:  make(chan struct{}),
	}
}

func (u *Updater) Start(ctx context.Context) error {
	var (
		localCancel context.CancelFunc
		localCtx    context.Context
	)

	u.timer = time.NewTicker(time.Duration(u.config.IntervalMinutes) * time.Minute)

	cancel := func() {
		if localCancel != nil {
			localCancel()
		}
	}
	run := func() {
		localCtx, localCancel = context.WithCancel(ctx)
		go u.run(localCtx)
	}

	go func() {
		run() // first time
		for {
			select {
			case <-u.timer.C:
				cancel()
				run()
			case <-ctx.Done():
				cancel()
				return

			case <-u.stop:
				cancel()
				return
			}
		}
	}()
	return nil
}

func (u *Updater) Stop(_ context.Context) error {
	if u.stop != nil {
		u.stop <- struct{}{}
	}
	return nil
}

func (u *Updater) run(ctx context.Context) error {
	log.Println("running updater")

	log.Println("fetching remote timestamp")
	remote, err := u.remote.LatestData(ctx)
	if err != nil {
		return err
	}

	log.Println("fetching local timestamp")
	local, err := u.sync.CurrentDate(ctx)
	if err != nil {
		return err
	}

	if local.Equal(remote.Timestamp) {
		log.Println("no updates. finished updater")
		return nil
	} else {
		log.Printf("new timestamp with value %v\n", remote.Timestamp)
	}

	log.Println("downloading latest data")
	zipRaw, err := u.remote.DownloadLatestData(ctx, remote)
	if err != nil {
		return err
	}

	zip, err := u.file.FromRaw(zipRaw)
	if err != nil {
		return err
	}

	log.Println("extracting zip")
	file, err := u.file.ExtractZip(
		zip,
		func(s string) bool { return strings.HasSuffix(s, ".sql") },
	)
	if err != nil {
		return err
	}

	log.Println("import sql file")
	if err := u.sync.ImportFile(file); err != nil {
		return err
	}

	log.Println("updating database")
	if err := u.sync.Update(ctx); err != nil {
		return err
	}

	log.Println("refreshing database")
	return u.sync.Refresh(ctx)
}
