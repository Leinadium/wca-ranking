package wca

import "leinadium.dev/wca-ranking/internal/adapter/config"

func NewWCATokenService(config config.WCA) *WCATokenService {
	return &WCATokenService{
		requester:         NewRequester(),
		clientID:          config.ClientId,
		clientSecret:      config.ClientSecret,
		endpointAuthorize: config.Endpoints.OAuthAuthorize,
		endpointToken:     config.Endpoints.OAuthToken,
		redirectURI:       config.RedirectURI,
	}
}

type WCATokenService struct {
	requester *Requester

	clientID          string
	clientSecret      string
	endpointAuthorize string
	endpointToken     string
	redirectURI       string
}
