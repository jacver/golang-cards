package main

import (
	"os"
	"testing"
)

// All test filenames end with _test.go
// All test files run with >go test in terminal

// Determining what to test:
// Identify easy assertions that should be true if function is working

// In regard to newDeck()
// // Check length of deck since we know what it should be
// // First card should be ace of spades
// // Last card should be four of clubs
// // A more test would be testing the file save and file load
// // // Note on this: If the file has a problem after creation, it also created the test file. This must MANUALLY be deleted in testing
// // // To ensure testing in clean environment, delete all files w/ name on start and on end

func TestNewDeck(t *testing.T) {
	// t is test handler

	d := newDeck()

	// testing deck length
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// Testing order of cards
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting") // remove old test files

	deck := newDeck() // create new deck
	deck.saveToFile("_decktesting") // save deck to file

	loadedDeck := newDeckFromFile("_decktesting") // load deck from file

	// assertion 
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting") // remove test file
}