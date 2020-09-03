package main

func main() {
	cards := newDeck()

	cards.saveToFile("my_cards")

	//fmt.Println(cards.toString())

	//hand, restOfCards := deal(cards, 2)
	//hand.print()
	//restOfCards.print()
}
