package domain

import "time"

type User struct {
	WCAID        WCAID
	StateID      string
	RegisterDate time.Time
}

type UserStatus struct {
	WCAID        WCAID
	IsRegistered bool
	CanRegister  bool
	Updated      time.Time
}

type UserBasic struct {
	WCAID   WCAID
	Name    string
	Country string
}
