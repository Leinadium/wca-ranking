package auth

import (
	"context"
	"net/url"
	"time"

	"leinadium.dev/wca-ranking/internal/adapter/config"
	"leinadium.dev/wca-ranking/internal/adapter/request"
	"leinadium.dev/wca-ranking/internal/core/domain"
)

func NewAuthRequester(config *config.WCA) *AuthRequester {
	return &AuthRequester{
		requester: request.NewRequester(),
		config:    config,
	}
}

type AuthRequester struct {
	requester *request.Requester
	config    *config.WCA
}

func (r *AuthRequester) AccessToken(ctx context.Context, code string) (*domain.AuthPayload, error) {
	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("client_id", r.config.ClientID)
	values.Add("client_secret", r.config.ClientSecret)
	values.Add("redirect_uri", r.config.RedirectURI)

	var res oAuthResponse
	err := r.requester.PostJSON(
		r.config.Endpoints.OAuthToken,
		values,
		&res,
	)
	if err != nil {
		return nil, err
	} else {
		return &domain.AuthPayload{
			AccessToken: res.AccessToken,
			ExpiresAt:   time.Now().Add(time.Second * time.Duration(res.ExpiresIn)),
		}, nil
	}
}

type oAuthResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	CreatedAt    int    `json:"created_at"`
}
