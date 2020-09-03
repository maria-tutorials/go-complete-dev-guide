package main

func main() {
	cards := newDeck()

	hand, restOfCards := deal(cards, 2)

	hand.print()
	restOfCards.print()
}
