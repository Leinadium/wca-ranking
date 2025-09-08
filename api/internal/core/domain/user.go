package domain

import "time"

type User struct {
	WCAID        string
	StateID      string
	RegisterDate time.Time
}

type UserStatus struct {
	WCAID        string
	IsRegistered bool
	CanRegister  bool
	Updated      time.Time
}

type UserBasic struct {
	WCAID   string
	Name    string
	Country string
}
