package main

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

func (m Match) HomeTeamScoreUpdate() {
	m.HomeTeamScore += 1
}

func (m Match) AwayTeamScoreUpdate() {
	m.AwayTeamScore += 1
}
