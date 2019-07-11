package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16{
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}

	// d.shuffle()

	if d[0] != "Ace of Spades"{
		t.Errorf("Expected first card in deck to `Ace of Spades` got %v",d[0])
	}

	if d[len(d) -1] != "Four of Clubs"{
		t.Errorf("Expected last card to be `Four of Clubs` got %v",d[len(d) -1])
	}
}


func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16{
		t.Errorf("Expected 16 cards in deck got %v",len(loadedDeck))
	}

	os.Remove("_decktesting")
}