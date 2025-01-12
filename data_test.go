package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mixOrder(t *testing.T) {
	type testCaseParams struct {
		inData  func() []string
		outData func() []string
	}

	type testCaseExpectedValues struct {
		orderChanged bool
	}

	type testCase struct {
		name           string
		params         testCaseParams
		expectedValues testCaseExpectedValues
	}

	testCases := []testCase{
		{
			name: "Order not mixed",
			params: testCaseParams{
				inData: func() []string {
					return teams[:]
				},
				outData: func() []string {
					return teams[:]
				},
			},
			expectedValues: testCaseExpectedValues{
				orderChanged: false,
			},
		},
		{
			name: "Order mixed",
			params: testCaseParams{
				inData: func() []string {
					return teams[:]
				},
				outData: func() []string {
					s := mixOrder(teams)
					return s[:]
				},
			},
			expectedValues: testCaseExpectedValues{
				orderChanged: true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expectedValues.orderChanged {
				require.False(t, slices.Equal(tc.params.inData(), tc.params.outData()))
			} else {
				require.True(t, slices.Equal(tc.params.inData(), tc.params.outData()))
			}
		})
	}
}
