package domain

import "github.com/guregu/null"

type WCAID string

// Person is the basic information of a person on the database
type Person struct {
	Name              string
	State             null.String
	Registered        bool
	TotalCompetitions int
	StateCompetitions null.Int
}
