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

// match is a routine representing a game between two teams.
func match(ctx context.Context, wg *sync.WaitGroup, eventChannel chan MatchEvent, m Match) {
	defer wg.Done()

	tickerTimeout := time.NewTicker(5 * time.Second)
	tickerEvent := time.NewTicker(time.Second)

	eventChannel <- MatchEvent{
		Type:  EventStartMatch,
		Match: m,
	}

	homeTeamScore, awayTeamScore := m.HomeTeamScore, m.AwayTeamScore

	var done bool
	for !done {
		select {
		case <-ctx.Done():
			done = true
		case <-tickerEvent.C:
			homeTeamScore += 1 // TODO - add some logic about random scoring by each team.
			eventChannel <- MatchEvent{
				Type: EventUpdateMatchScore,
				Match: Match{
					HomeTeam:      m.HomeTeam,
					AwayTeam:      m.AwayTeam,
					HomeTeamScore: homeTeamScore,
					AwayTeamScore: awayTeamScore,
					ReferenceID:   m.ReferenceID,
				},
			}
		case <-tickerTimeout.C:
			eventChannel <- MatchEvent{
				Type: EventStopMatch,
				Match: Match{
					HomeTeam:      m.HomeTeam,
					AwayTeam:      m.AwayTeam,
					HomeTeamScore: homeTeamScore,
					AwayTeamScore: awayTeamScore,
					ReferenceID:   m.ReferenceID,
				},
			}
			done = true
		}
	}
}

// board is a routine representing score board collecting data of the played matches.
func board(ctx context.Context, cf context.CancelFunc, wg *sync.WaitGroup, eventChannel chan MatchEvent) {
	defer wg.Done()

	tickerTimeout := time.NewTicker(5 * time.Second)

	var done bool
	for !done {
		select {
		case <-ctx.Done():
			done = true
		case e, _ := <-eventChannel: // TODO - add recognizing if channel if closed.
			fmt.Printf("got event: %v\n", e)
		case <-tickerTimeout.C:
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
	waitGroup.Add(1)
	waitGroup.Add(1)

	eventChannel := make(chan MatchEvent, 64)

	m := Match{
		HomeTeam:      "Poland",
		AwayTeam:      "England",
		HomeTeamScore: 0,
		AwayTeamScore: 0,
		ReferenceID:   1,
	}
	m2 := Match{
		HomeTeam:      "Mexico",
		AwayTeam:      "Qatar",
		HomeTeamScore: 0,
		AwayTeamScore: 0,
		ReferenceID:   2,
	}

	go board(ctx, cancel, waitGroup, eventChannel)
	go match(ctx, waitGroup, eventChannel, m)
	go match(ctx, waitGroup, eventChannel, m2)

	<-ctx.Done()

	waitGroup.Wait()
}

func main() {
	ctx := context.Background()

	fmt.Println("Hello Sport Radar :)")

	run(ctx)
}
