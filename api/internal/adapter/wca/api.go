package wca

import (
	"context"
	"time"

	"leinadium.dev/wca-ranking/internal/adapter/wca/models"
	"leinadium.dev/wca-ranking/internal/core/domain"
)

const (
	dateFormat = "2006-01-02T15:04:05Z"
)

type WCAAPIRequester struct {
	requester          *Requester
	latestDataEndpoint string
	userEndpoint       string
}

func (r *WCAAPIRequester) LatestData(ctx context.Context) (*domain.WCALatestData, error) {
	var payload models.WCALatestData
	if err := r.requester.GetJSON(r.latestDataEndpoint, &payload); err != nil {
		return nil, err
	}

	res, err := time.Parse(dateFormat, payload.ExportDate)
	if err != nil {
		return nil, err
	}

	return &domain.WCALatestData{Timestamp: res, DownloadUrl: payload.SqlUrl}, nil
}

func (r *WCAAPIRequester) DownloadLatestData(ctx context.Context, data *domain.WCALatestData) (domain.RawFile, error) {
	body, err := r.requester.Get(data.DownloadUrl)
	if err != nil {
		return nil, err
	}
	return domain.RawFile(body), nil
}

func (r *WCAAPIRequester) UserInfo(ctx context.Context, accessToken string) (*domain.UserBasic, error) {
	var payload models.WCAUserInfo
	if err := r.requester.GetJSONAuthenticated(r.userEndpoint, accessToken, &payload); err != nil {
		return nil, err
	}
	return &domain.UserBasic{
		WCAID:   payload.Me.WCAID,
		Name:    payload.Me.Name,
		Country: payload.Me.Country,
	}, nil
}
