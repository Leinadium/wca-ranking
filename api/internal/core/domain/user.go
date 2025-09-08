package domain

import "time"

type User struct {
	WCAID        string
	StateID      string
	RegisterDate time.Time
}
