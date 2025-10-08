package user

import (
	"context"

	"leinadium.dev/wca-ranking/internal/adapter/config"
	"leinadium.dev/wca-ranking/internal/adapter/request"
	"leinadium.dev/wca-ranking/internal/core/domain"
)

func NewUserRequester(config *config.WCA) *UserRequester {
	return &UserRequester{
		requester: request.NewRequester(),
		config:    config,
	}
}

type UserRequester struct {
	requester *request.Requester
	config    *config.WCA
}

func (r *UserRequester) UserInfo(ctx context.Context, accessToken string) (*domain.UserBasic, error) {
	var payload userInfo
	if err := r.requester.GetJSONAuthenticated(r.config.Endpoints.Me, accessToken, &payload); err != nil {
		return nil, err
	}
	return &domain.UserBasic{
		WCAID:   domain.WCAID(payload.Me.WCAID),
		Name:    payload.Me.Name,
		Country: payload.Me.Country,
	}, nil
}

type userInfo struct {
	Me struct {
		Name    string `json:"name"`
		WCAID   string `json:"wca_id"`
		Country string `json:"country_iso2"`
	} `json:"me"`
}
