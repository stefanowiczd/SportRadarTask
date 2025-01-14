//go:build unit

package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ScoreBoard(t *testing.T) {
	type testCaseParams struct {
		inData  func() []Match
		outData func() []Match
	}

	type testCaseExpectedValues struct{}

	type testCase struct {
		name           string
		params         testCaseParams
		expectedValues testCaseExpectedValues
	}

	testCases := []testCase{
		{
			name: "Scenario 0",
			params: testCaseParams{
				inData: func() []Match {
					return []Match{
						{
							HomeTeam:      "Mexico",
							AwayTeam:      "Canada",
							HomeTeamScore: 0,
							AwayTeamScore: 5,
							ReferenceID:   1,
						},
						{
							HomeTeam:      "Spain",
							AwayTeam:      "Brazil",
							HomeTeamScore: 10,
							AwayTeamScore: 2,
							ReferenceID:   2,
						},
						{
							HomeTeam:      "Germany",
							AwayTeam:      "France",
							HomeTeamScore: 2,
							AwayTeamScore: 2,
							ReferenceID:   3,
						},
						{
							HomeTeam:      "Uruguay",
							AwayTeam:      "Italy",
							HomeTeamScore: 6,
							AwayTeamScore: 6,
							ReferenceID:   4,
						},
						{
							HomeTeam:      "Argentina",
							AwayTeam:      "Australia",
							HomeTeamScore: 3,
							AwayTeamScore: 1,
							ReferenceID:   5,
						},
					}
				},
				outData: func() []Match {
					return []Match{
						{
							HomeTeam:      "Uruguay",
							AwayTeam:      "Italy",
							HomeTeamScore: 6,
							AwayTeamScore: 6,
							ReferenceID:   4,
						},
						{
							HomeTeam:      "Spain",
							AwayTeam:      "Brazil",
							HomeTeamScore: 10,
							AwayTeamScore: 2,
							ReferenceID:   2,
						},
						{
							HomeTeam:      "Mexico",
							AwayTeam:      "Canada",
							HomeTeamScore: 0,
							AwayTeamScore: 5,
							ReferenceID:   1,
						},
						{
							HomeTeam:      "Argentina",
							AwayTeam:      "Australia",
							HomeTeamScore: 3,
							AwayTeamScore: 1,
							ReferenceID:   5,
						},
						{
							HomeTeam:      "Germany",
							AwayTeam:      "France",
							HomeTeamScore: 2,
							AwayTeamScore: 2,
							ReferenceID:   3,
						},
					}
				},
			},
		},
		{
			name: "Scenario 1",
			params: testCaseParams{
				inData: func() []Match {
					return []Match{
						{
							HomeTeam:      "Poland",
							AwayTeam:      "England",
							HomeTeamScore: 0,
							AwayTeamScore: 0,
							ReferenceID:   1,
						},
						{
							HomeTeam:      "Germany",
							AwayTeam:      "France",
							HomeTeamScore: 4,
							AwayTeamScore: 0,
							ReferenceID:   2,
						},
						{
							HomeTeam:      "Spain",
							AwayTeam:      "USA",
							HomeTeamScore: 1,
							AwayTeamScore: 1,
							ReferenceID:   3,
						},
					}
				},
				outData: func() []Match {
					return []Match{
						{
							HomeTeam:      "Germany",
							AwayTeam:      "France",
							HomeTeamScore: 4,
							AwayTeamScore: 0,
							ReferenceID:   2,
						},
						{
							HomeTeam:      "Spain",
							AwayTeam:      "USA",
							HomeTeamScore: 1,
							AwayTeamScore: 1,
							ReferenceID:   3,
						},
						{
							HomeTeam:      "Poland",
							AwayTeam:      "England",
							HomeTeamScore: 0,
							AwayTeamScore: 0,
							ReferenceID:   1,
						},
					}
				},
			},
		},
		{
			name: "Scenario 2",
			params: testCaseParams{
				inData: func() []Match {
					return []Match{
						{
							HomeTeam:      "Poland",
							AwayTeam:      "England",
							HomeTeamScore: 0,
							AwayTeamScore: 0,
							ReferenceID:   0,
						},
						{
							HomeTeam:      "Germany",
							AwayTeam:      "France",
							HomeTeamScore: 4,
							AwayTeamScore: 0,
							ReferenceID:   2,
						},
						{
							HomeTeam:      "Spain",
							AwayTeam:      "USA",
							HomeTeamScore: 1,
							AwayTeamScore: 1,
							ReferenceID:   3,
						},
						{
							HomeTeam:      "Argentina",
							AwayTeam:      "Mexico",
							HomeTeamScore: 1,
							AwayTeamScore: 6,
							ReferenceID:   4,
						},
						{
							HomeTeam:      "Denmark",
							AwayTeam:      "Japan",
							HomeTeamScore: 2,
							AwayTeamScore: 0,
							ReferenceID:   5,
						},
					}
				},
				outData: func() []Match {
					return []Match{
						{
							HomeTeam:      "Argentina",
							AwayTeam:      "Mexico",
							HomeTeamScore: 1,
							AwayTeamScore: 6,
							ReferenceID:   4,
						},
						{
							HomeTeam:      "Germany",
							AwayTeam:      "France",
							HomeTeamScore: 4,
							AwayTeamScore: 0,
							ReferenceID:   2,
						},
						{
							HomeTeam:      "Denmark",
							AwayTeam:      "Japan",
							HomeTeamScore: 2,
							AwayTeamScore: 0,
							ReferenceID:   5,
						},
						{
							HomeTeam:      "Spain",
							AwayTeam:      "USA",
							HomeTeamScore: 1,
							AwayTeamScore: 1,
							ReferenceID:   3,
						},
						{
							HomeTeam:      "Poland",
							AwayTeam:      "England",
							HomeTeamScore: 0,
							AwayTeamScore: 0,
							ReferenceID:   0,
						},
					}
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			sb := NewScoreBoard()

			matches := tc.params.inData()
			for i := 0; i < len(matches); i++ {
				sb.StartMatch(matches[i])
				require.NoError(t, sb.UpdateMatchScore(matches[i]))
				sb.StopMatch(matches[i])
			}

			sb.SortResult()

			require.True(t, slices.Equal(sb.Results, tc.params.outData()))
		})
	}

}
