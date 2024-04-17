package main

func main() {
	cards := deck{"Ace of Spades", newcard()} // This is equivalent to []string{}
	cards.print()
}

// go wants us to always add what type of values a function would return
func newcard() string {
	return "Five of Diamonds"
}
