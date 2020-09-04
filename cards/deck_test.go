package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 6 {
		t.Errorf("Expected deck lenght of 6, but got %v", len(d))
	}

	if d[0] != "Ace of Diamonds" {
		t.Errorf("Expected first cards to be Ace of Diamonds, but got %v", d[0])
	}

	if d[len(d)-1] != "Three of Hearts" {
		t.Errorf("Expected last card to be Three of Hearts, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 6 {
		t.Errorf("Expected 6 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
