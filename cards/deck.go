package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
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

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Creating the deck...")
		return newDeck()
		//os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		newpos := rand.Intn(len(d) - 1)

		d[i], d[newpos] = d[newpos], d[i] //weird-magic swap
	}
}
