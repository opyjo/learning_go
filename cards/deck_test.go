package main

import "testing"

func TestNewDeck(t *testing.T) {
	// Create a new deck
	d := newDeck()

	// Check if the deck has the correct number of cards
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// Check if the first card is an Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	// Check if the last card is a Four of Clubs
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}

}
