package domain

import "time"

type AuthPayload struct {
	AccessToken string
	ExpiresAt   time.Time
}
