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

	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		time.Sleep(5 * time.Second)

		cancel()

	}(waitGroup)

	<-ctx.Done()
}

func main() {
	ctx := context.Background()

	fmt.Println("Hello Sport Radar :)")

	run(ctx)
}
