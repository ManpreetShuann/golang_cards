package main

import "fmt"

func main() {
	cards := deck{"Ace of Spades", newcard()}
	for _, card := range cards {
		fmt.Println(card)
	}
}

// go wants us to always add what type of values a function would return
func newcard() string {
	return "Five of Diamonds"
}
