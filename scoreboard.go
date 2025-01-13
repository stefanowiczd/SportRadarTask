package main

import (
	"errors"
	"fmt"
	"sync"
)

var errItemNotFound = errors.New("item not found")

type ScoreBoard struct {
	OngoingGames sync.Map
	Results      []Match
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
