package main

import (
	"os"
	"testing"
)

//t represents test handler
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 20 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected First card of Ace of spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last of four of clubs but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFine(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 116 {
		t.Errorf("Expected 16 cards but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
