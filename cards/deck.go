package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardTypes := []string{"Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, ctype := range cardTypes {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+ctype)
		}
	}

	return cards
}

//receiver function:
// any variable of type deck gets access to print
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
