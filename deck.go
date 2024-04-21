package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Creates a new type of deck which is a slice of strings.
type deck []string

// Creates and returns a New deck of cards.
func newDeck() deck {
	cards := deck{}
	cardsSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardsValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}
	for _, suit := range cardsSuits {
		for _, value := range cardsValues {
			// cards = append(cards, value+" of "+suit)
			cards = append(cards, fmt.Sprintf("%s of %s", value, suit))
		}
	}
	return cards
}

// This is a reciever function to print all cards in the deck.
// We added a reciever(d deck) which means any variable of type deck now gets access to print method, Kinda similar to self in python.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// A reciever function count, to count the number of cards in the deck.
func (d deck) count() {
	count := 0
	for range d {
		count += 1
	}
	fmt.Println(count)
}

// A parameterized function to deal a hand of cards
func deal(d deck, handSize int) (deck, deck) {
	hand := d[:handSize]
	deck := d[handSize:]
	return hand, deck
}

// A function to convert a deck to a string
func (d deck) toString() string {
	deckString := strings.Join(d, ",")
	return deckString
}

// Save a deck to file
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// Creates a new deck from a file
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error: Couldn't open file %s: %v", filename, err)
	}
	s := string(bs)
	d := strings.Split(s, ",")
	return d
}
