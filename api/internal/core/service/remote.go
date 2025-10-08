package service

import (
	"context"

	"leinadium.dev/wca-ranking/internal/core/domain"
	"leinadium.dev/wca-ranking/internal/core/port"
)

type RemoteService struct {
	r port.RemoteRequester
}

func NewRemoteService(r port.RemoteRequester) *RemoteService {
	return &RemoteService{r: r}
}

func (r *RemoteService) LatestData(ctx context.Context) (*domain.RemoteLatestData, error) {
	return r.r.LatestData(ctx)
}

func (r *RemoteService) DownloadLatestData(ctx context.Context, data *domain.RemoteLatestData) (domain.RawFile, error) {
	return r.r.DownloadLatestData(ctx, data)
}
