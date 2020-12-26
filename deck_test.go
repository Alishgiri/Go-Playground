package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	dlen := len(d)

	if dlen != 12 {
		t.Errorf("Expected deck length of 12, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[dlen-1] != "Three of Clubs" {
		t.Errorf("Expected first card of Three of Clubs, but got %v", d[dlen-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testFileName := "_decttesting"
	os.Remove(testFileName)

	deck := newDeck()
	deck.saveToFile(testFileName)

	lodedDeck := newDeckFromFile(testFileName)

	if len(lodedDeck) != 12 {
		t.Errorf("Expected 16 cards in deck, got %v", len(lodedDeck))
	}

	os.Remove(testFileName)
}
