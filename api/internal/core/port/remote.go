package port

import (
	"context"

	"leinadium.dev/wca-ranking/internal/core/domain"
)

type RemoteRequester interface {
	// GetLatestData obtains the latest data using WCA's API
	LatestData(ctx context.Context) (*domain.RemoteLatestData, error)

	// DownloadLatestData downloads the latest data into a file
	DownloadLatestData(ctx context.Context, data *domain.RemoteLatestData) (domain.RawFile, error)
}

type RemoteService interface {
	// GetLatestData obtains the latest data using WCA's API
	LatestData(ctx context.Context) (*domain.RemoteLatestData, error)

	// DownloadLatestData downloads the latest available data into a file
	DownloadLatestData(ctx context.Context, data *domain.RemoteLatestData) (domain.RawFile, error)
}
