package simulation

import (
	"context"
	"fmt"
	"log"
	"manabase-simulation/package/logging"
	"manabase-simulation/package/model"
	"net/http"
	"sync"
	"time"
)

const (
	// CheckpointInterval is the number of iterations between checkpoints.
	CheckpointInterval = 100

	// TotalCheckpoints is the total number of checkpoints to generate.
	TotalCheckpoints = 30
)

// SimulateDeck Simulates a deck against a given objective with the provided configuration.
func SimulateDeck(ctx context.Context, deckList model.DeckList, gameConfiguration model.GameConfiguration, objective model.TestObjective) bool {
	//logger := CreateLogger()
	//logger.Debug("Starting deck simulation")

	// Generate Randomized Deck
	deck := deckList.GenerateDeck()
	hand := model.NewDeck()
	board := NewBoardState()

	// TODO: Add validations like Validate deck is >= 60 cards

	// Draw Initial Hand
	for range gameConfiguration.InitialHandSize {
		hand = deck.DrawCard(hand)
	}

	// For turnNumber to target turn
	for turnNumber := range objective.TargetTurn {
		// If turnNumber = 1 and on the play, skip draw
		if turnNumber == 0 && gameConfiguration.OnThePlay {
			// Skip your draw
			//logger.Debug("Playing first, skipping draw")
		} else {
			hand = deck.DrawCard(hand)
		}

		hand = board.PlayLand(hand, objective, turnNumber+1)
	}

	// Compute if target is met (possibly using backtracking?)
	// Computation can start with the most restrictive lands by sorting based on number of colors it taps for.
	isMet, _ := board.ValidateTestObjective(objective)
	return isMet
}

func Simulate(ctx context.Context, decklist model.DeckList, configuration model.GameConfiguration, objective model.TestObjective) []model.ResultCheckpoint {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	logger := logging.CreateLogger()
	logger.Info(decklist.ToString())
	logger.Info(objective.ToString())

	now := time.Now()

	successCount := 0
	iterations := CheckpointInterval * TotalCheckpoints

	resultChannel := make(chan bool, 100)
	wg := new(sync.WaitGroup)

	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go StartSimulation(ctx, decklist, configuration, objective, resultChannel, wg)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	completedIterations := 0
	checkpoints := make([]model.ResultCheckpoint, TotalCheckpoints)
	for result := range resultChannel {
		completedIterations++
		if result {
			successCount++
		}

		if completedIterations%CheckpointInterval == 0 {
			checkpoints[(completedIterations/CheckpointInterval)-1] = model.ResultCheckpoint{
				Iterations: int32(completedIterations),
				Successes:  int32(successCount),
			}
		}

	}

	// Capture results to be consumes.
	logger.Info(fmt.Sprintf("Success count: %d", checkpoints[len(checkpoints)-1].Successes))
	logger.Info(fmt.Sprintf("Success Rate: %f", checkpoints[len(checkpoints)-1].GetSuccessRate()))
	logger.Info(fmt.Sprintf("Checkpoints: %v", checkpoints))
	logger.Info(fmt.Sprintf("Time taken: %s", time.Since(now)))

	return checkpoints
}

func StartSimulation(ctx context.Context, deckList model.DeckList, gameConfiguration model.GameConfiguration, objective model.TestObjective, c chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	c <- SimulateDeck(ctx, deckList, gameConfiguration, objective)
}
