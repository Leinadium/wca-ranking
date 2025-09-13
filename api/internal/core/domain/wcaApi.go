package domain

import "time"

type WCALatestData struct {
	Timestamp   time.Time
	DownloadUrl string
}
