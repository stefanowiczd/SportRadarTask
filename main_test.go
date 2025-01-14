//go:build race_cond

package main

import (
	"context"
	"testing"
)

func Test_run(t *testing.T) {
	t.Run("Race condition hazard", func(t *testing.T) {
		ctx := context.Background()

		run(ctx)
	})
}
