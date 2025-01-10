package main

const (
	EventUnknown          = 0
	EventStartMatch       = 1
	EventStopMatch        = 2
	EventUpdateMatchScore = 3
)

// MatchEvent represents a match event, like start match, update score or finish game.
type MatchEvent struct {
	Type int
	Match
}

type Match struct {
	HomeTeam      string
	AwayTeam      string
	HomeTeamScore int
	AwayTeamScore int
	ReferenceID   int
}

func NewMatch(ht, at string, hts, ats, id int) Match {
	return Match{
		HomeTeam:      ht,
		AwayTeam:      at,
		HomeTeamScore: hts,
		AwayTeamScore: ats,
		ReferenceID:   id,
	}
}
