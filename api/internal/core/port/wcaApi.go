package port

import (
	"context"

	"leinadium.dev/wca-ranking/internal/core/domain"
)

type WCAAPIRequester interface {
	// GetLatestData obtains the latest data using WCA's API
	LatestData(ctx context.Context) (*domain.WCALatestData, error)

	// DownloadLatestData downloads the latest data into a file
	DownloadLatestData(ctx context.Context, data *domain.WCALatestData) (domain.RawFile, error)

	// UserInfo obtains the current basic information from WCA's API
	UserInfo(ctx context.Context, accessToken string) (*domain.UserBasic, error)
}

type WCAAPIService interface {
	// GetLatestData obtains the latest data using WCA's API
	LatestData(ctx context.Context) (*domain.WCALatestData, error)

	// DownloadLatestData downloads the latest available data into a file
	DownloadLatestData(ctx context.Context, data *domain.WCALatestData) (domain.File, error)

	// UserInfo obtains the current basic information from WCA's API
	UserInfo(ctx context.Context, accessToken string) (*domain.UserBasic, error)
}
