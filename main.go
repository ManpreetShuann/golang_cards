package main

// package main tells us that this package is an executable

func main() {
	cards := newDeck()
	// cards := newDeckFromFile("my_hand")
	// cards.print()
	cards.count()
	handSize := 5
	hand, remainingCards := deal(cards, handSize)
	hand.print()
	remainingCards.print()
	hand.toString()
	hand.saveToFile("../my_hand")
}
