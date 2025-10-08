package remote

import (
	"context"
	"time"

	"leinadium.dev/wca-ranking/internal/adapter/config"
	"leinadium.dev/wca-ranking/internal/adapter/request"
	"leinadium.dev/wca-ranking/internal/core/domain"
)

const (
	dateFormat = "2006-01-02T15:04:05Z"
)

func NewRemoteRequester(config *config.WCA) *RemoteRequester {
	return &RemoteRequester{
		requester: request.NewRequester(),
		config:    config,
	}
}

type RemoteRequester struct {
	requester *request.Requester
	config    *config.WCA
}

func (r *RemoteRequester) LatestData(ctx context.Context) (*domain.RemoteLatestData, error) {
	var payload latestData
	if err := r.requester.GetJSON(r.config.Endpoints.LatestData, &payload); err != nil {
		return nil, err
	}

	res, err := time.Parse(dateFormat, payload.ExportDate)
	if err != nil {
		return nil, err
	}

	return &domain.RemoteLatestData{Timestamp: res, DownloadUrl: payload.SqlUrl}, nil
}

func (r *RemoteRequester) DownloadLatestData(ctx context.Context, data *domain.RemoteLatestData) (domain.RawFile, error) {
	body, err := r.requester.Get(data.DownloadUrl)
	if err != nil {
		return nil, err
	}
	return domain.RawFile(body), nil
}

type latestData struct {
	ExportDate string `json:"export_date"`
	SqlUrl     string `json:"sql_url"`
	TsvUrl     string `json:"tsv_url"`
}
