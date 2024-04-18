package main

import "fmt"

// Create a new type of deck which is a slice of strings
type deck []string

// Creates and returns a New deck of cards
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

// reciever function: We added a reciever(d deck).
// which means any variable of type deck now gets access to print method, Kinda similar to self in python.
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) count() {
	count := 0
	for range d {
		count += 1
	}
	fmt.Println(count)
}
