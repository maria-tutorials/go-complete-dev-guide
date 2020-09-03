package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())

	//hand, restOfCards := deal(cards, 2)
	//hand.print()
	//restOfCards.print()
}
