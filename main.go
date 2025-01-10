package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func scheduler(ctx context.Context, cf context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done()

	var done bool
	for !done {
		select {
		case <-ctx.Done():
			done = true
		case <-time.After(time.Second * 3):
			done = true
		}
	}

	cf()
}

func run(ctx context.Context) {
	waitGroup := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(ctx)

	// Dedicated loop for catching system signals.
	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

		sig := <-signalChan

		slog.WarnContext(ctx, "Received signal", slog.String("signal", sig.String()))
		cancel()
	}()

	waitGroup.Add(1)

	go scheduler(ctx, cancel, waitGroup)

	<-ctx.Done()

	waitGroup.Wait()
}

func main() {
	ctx := context.Background()

	fmt.Println("Hello Sport Radar :)")

	run(ctx)
}
