package domain

import "time"

type RemoteLatestData struct {
	Timestamp   time.Time
	DownloadUrl string
}
