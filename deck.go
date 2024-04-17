package main

import "fmt"

// Create a new type of deck which is a slice of strings
type deck []string

// reciever function: We added a reciever(d deck), which means any variable of type deck now gets access to print method
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
