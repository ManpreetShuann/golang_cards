package main

// package main tells us that this package is an executable

import "fmt"

func main() {
	cards := newDeck()
	// cards := newDeckFromFile("my_hand")
	cards.shuffle()
	// cards.print()
	handSize := 4
	hand, remainingCards := deal(cards, handSize)
	hand.print()
	fmt.Println("Remaining Cards in Deck:", remainingCards.count())
	// hand.toString()
	// hand.saveToFile("../my_hand")
}
