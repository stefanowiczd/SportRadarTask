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

	// Two timeouts to simulate of played match
	tickerTimeout := time.NewTicker(6 * time.Second) // How long it will take
	tickerEvent := time.NewTicker(time.Second)       // When the decision about action will be taken (scored goal or no goal)

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
			// Pseudo logic to fake action of scoring a goal by home or away team, or no goal.
			if randRange(0, 4)%3 == 0 { // No goal was scored, do not send event.
				continue
			} else if randRange(0, 4)%3 == 1 {
				homeTeamScore += 1 // Goal scored by home team.
			} else {
				awayTeamScore += 1 // Goal scored by away team.
			}

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

	scoreBoard := NewScoreBoard()
	tickerTimeout := time.NewTicker(7 * time.Second)

	var done bool
	for !done {
		select {
		case <-ctx.Done():
			done = true
		case e, _ := <-eventChannel: // TODO - add recognizing if channel if closed.
			switch e.Type {
			case EventStartMatch:
				scoreBoard.StartMatch(e.Match)
			case EventUpdateMatchScore:
				if err := scoreBoard.UpdateMatchScore(e.Match); err != nil {
					slog.ErrorContext(ctx, "Can not process received event", slog.Any("error", err), slog.Any("match", e.Match))
				}
			case EventStopMatch:
				scoreBoard.StopMatch(e.Match)
			default:
				slog.ErrorContext(ctx, "Unknown event", slog.Int("event_value", e.Type))
			}
		case <-tickerTimeout.C:
			done = true
		}
	}

	scoreBoard.SortResult() // Sort result by the sum of scored goals.

	scoreBoard.Summary()

	// Only board should be able to realize cancel function since it has to wait until finish of all matches.
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
	eventChannel := make(chan MatchEvent, 64)
	go board(ctx, cancel, waitGroup, eventChannel)

	teamsWorldCup := mixOrder(teams)

	for i, j := 0, 0; i <= len(teamsWorldCup)-1; i, j = i+2, j+1 {
		m := Match{
			HomeTeam:    teamsWorldCup[i],
			AwayTeam:    teamsWorldCup[i+1],
			ReferenceID: j,
		}

		waitGroup.Add(1)

		go match(ctx, waitGroup, eventChannel, m)
	}

	<-ctx.Done()

	waitGroup.Wait()
}

func main() {
	ctx := context.Background()

	fmt.Printf("\nHello Sport Radar :)\n")

	fmt.Printf("\nThe tournament has just started...\n\n")

	run(ctx)
}
