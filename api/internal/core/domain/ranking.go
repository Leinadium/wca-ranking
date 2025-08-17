package domain

import (
	"github.com/guregu/null"
)

type RankingMode string
type EventID string

const (
	RankingAverage RankingMode = "average"
	RankingSingle  RankingMode = "single"
)

const (
	Event333           EventID = "333"
	Event222           EventID = "222"
	Event444           EventID = "444"
	Event555           EventID = "555"
	Event666           EventID = "666"
	Event777           EventID = "777"
	Event333Blind      EventID = "333bf"
	EventFMC           EventID = "333fm"
	EventFeet          EventID = "333ft"
	Event333OH         EventID = "333oh"
	EventClock         EventID = "clock"
	EventMegaminx      EventID = "minx"
	EventPyraminx      EventID = "pyram"
	EventSkewb         EventID = "skewb"
	EventSquare1       EventID = "sq1"
	Event444Blind      EventID = "444bf"
	Event555Blind      EventID = "555bf"
	Event333MultiBlind EventID = "333mbf"
)

// PersonRanking is an entry on the rankings of the person
type PersonRanking struct {
	WCAID            WCAID
	Mode             RankingMode
	Event            string
	Ranking          int
	Best             null.Int
	CompetitionId    string      // TODO: check if necessary
	CompetitionName  string      // TODO: check if necessary
	CompetitionState null.String // TODO: check if necessary
	Times            [5]null.Int
}

type PersonResult struct {
	Event          string
	Single         null.Int
	Average        null.Int
	RankingSingle  int
	RankingAverage int
}

type RankingEntry struct {
	WCAID            WCAID
	Name             string
	StateID          StateID
	Best             null.Int
	Ranking          int
	Registered       bool
	CompetitionID    string
	CompetitionName  string
	CompetitionState string
	Times            [5]null.Int
}
