package main

import (
	"fmt"
	"strings"
)

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

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

// aux function to turn deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
