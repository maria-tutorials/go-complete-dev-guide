package main

func main() {
	cards := newDeckFromFile("not_my_cards")

	cards.print()
	//	cards.saveToFile("my_cards")

}
