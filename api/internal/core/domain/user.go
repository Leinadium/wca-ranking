package domain

import "time"

type User struct {
	WcaID        string
	StateID      string
	RegisterDate time.Time
}
