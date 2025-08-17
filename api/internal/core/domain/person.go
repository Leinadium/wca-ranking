package domain

type WCAID string

// Person is the basic information of a person on the database
type Person struct {
	WCAID             WCAID
	Name              string
	State             string
	Registered        bool
	TotalCompetitions uint32
	StateCompetitions uint32
}
