package main

import (
	"fmt"

	"github.com/jzuhusky/gophercises/deck"
)

func main() {

	d := deck.NewDeck()
	for _, card := range d {
		fmt.Println(card)
	}
	fmt.Println("################################")
	deck.ShuffleDeck(d)
	for _, card := range d {
		fmt.Println(card)
	}
	fmt.Println("################################")
	deck.SortDeckSimple(d)
	for _, card := range d {
		fmt.Println(card)
	}
}
