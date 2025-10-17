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

	ticker *time.Ticker
	stop   chan struct{}
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

		ticker: nil,
		stop:   make(chan struct{}),
	}
}

func (u *Updater) Start(ctx context.Context) error {
	var (
		localCancel context.CancelFunc
		localCtx    context.Context
	)

	u.ticker = time.NewTicker(time.Duration(u.config.IntervalMinutes) * time.Minute)

	cancel := func() {
		if localCancel != nil {
			localCancel()
		}
	}
	run := func() {
		localCtx, localCancel = context.WithTimeout(context.Background(), time.Duration(5)*time.Minute)
		go u.runCatching(localCtx)
	}

	go func() {
		run() // first time
		for {
			select {
			case <-u.ticker.C:
				cancel()
				run()
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

func (u *Updater) runCatching(ctx context.Context) {
	if err := u.run(ctx); err != nil {
		log.Printf("could not run updater: %v\n", err)
	}
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

	log.Printf("local timestamp: %v\n", local)
	log.Printf("remote timestamp: %v\n", remote.Timestamp)

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

	log.Println("importing sql file")
	if err := u.sync.ImportFile(u.config.MariaDBBin, file); err != nil {
		return err
	}

	log.Println("updating database")
	if err := u.sync.Update(ctx); err != nil {
		return err
	}

	log.Println("refreshing database")
	if err := u.sync.Refresh(ctx); err != nil {
		return err
	}

	log.Printf("updating timestamp")
	if err := u.sync.SetCurrentDate(ctx, remote.Timestamp); err != nil {
		return err
	}

	log.Println("updater done")
	return nil
}
