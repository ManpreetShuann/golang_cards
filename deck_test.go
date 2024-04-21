package main

import "testing"

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 52 {
		t.Errorf("Expected 52 cards, got %d", len(cards))
	}
}
