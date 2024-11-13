package reader

import (
	"bufio"
	"fmt"
	"os"
)

const (
	DECK_START      = "Deck"
	SIDEBOARD_START = "Sideboard"
)

// ReadArenaDeckListFile reads an MTG Arena formatted decklist file and returns a decklist.
func ReadArenaDeckListFile(filename string) (deck []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	deckStarted := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == DECK_START {
			deckStarted = true
			continue
		}

		if line == SIDEBOARD_START {
			deckStarted = false
			break
		}

		if deckStarted {
			// Process the line here
			println(fmt.Sprintf("Adding %s to deck", line))
			if line != "" {
				deck = append(deck, line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return deck, nil
}
