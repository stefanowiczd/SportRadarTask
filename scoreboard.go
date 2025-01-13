package main

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

var errItemNotFound = errors.New("item not found")

type Results []Match

func (r Results) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Results) Len() int {
	return len(r)
}
func (r Results) Less(i, j int) bool {
	return r[i].HomeTeamScore+r[i].AwayTeamScore > r[j].HomeTeamScore+r[j].AwayTeamScore
}

type ScoreBoard struct {
	OngoingGames sync.Map
	Results
}

func NewScoreBoard() *ScoreBoard {
	return &ScoreBoard{
		OngoingGames: sync.Map{},
	}
}

func (s *ScoreBoard) StartMatch(m Match) {
	s.OngoingGames.Store(m.ReferenceID, m)
}

func (s *ScoreBoard) StopMatch(m Match) {
	s.OngoingGames.Delete(m.ReferenceID)

	s.Results = append(s.Results, m)
}

func (s *ScoreBoard) UpdateMatchScore(m Match) error {
	_, ok := s.OngoingGames.Swap(m.ReferenceID, m)
	if !ok {
		return fmt.Errorf("updating match score: %w", errItemNotFound)
	}

	return nil
}

// SortResult sorts result by the sum of scored goals.
func (s *ScoreBoard) SortResult() {
	sort.Sort(s.Results)
}

// Summary prints the result in user-friendly form.
func (s *ScoreBoard) Summary() {
	summary := s.Results
	for i := range summary {
		fmt.Printf(
			"%12s %d - %d %12s\n",
			summary[i].HomeTeam, summary[i].HomeTeamScore, summary[i].AwayTeamScore, summary[i].AwayTeam,
		)
	}

	fmt.Println()
}
