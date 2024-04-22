package main

import (
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 52 {
		t.Errorf("Expected 52 cards, got %d", len(cards))
	}
	spades := 0
	clubs := 0
	hearts := 0
	diamonds := 0
	for _, card := range cards {
		if strings.Contains(card, "Spades") {
			spades += 1
		}
		if strings.Contains(card, "Clubs") {
			clubs += 1
		}
		if strings.Contains(card, "Clubs") {
			hearts += 1
		}
		if strings.Contains(card, "Clubs") {
			diamonds += 1
		}
	}
	if spades != 13 {
		t.Errorf("Expected 13 Spades, got %d", spades)
	}
	if clubs != 13 {
		t.Errorf("Expected 13 Clubs, got %d", clubs)
	}
	if hearts != 13 {
		t.Errorf("Expected 13 Hearts, got %d", hearts)
	}
	if diamonds != 13 {
		t.Errorf("Expected 13 Diamonds, got %d", diamonds)
	}
}
