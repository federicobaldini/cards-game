package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0].value != "Ace" && d[0].suit != "Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1].value != "King" && d[len(d)-1].suit != "Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deckToSave := newDeck()
	deckToSave.saveToFile("_decktesting")

	deckToLoad := newDeckFromFile("_decktesting")

	if len(deckToLoad) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(deckToLoad))
	}

	os.Remove("_decktesting")
}
