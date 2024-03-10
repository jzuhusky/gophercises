package main

import (
    "fmt"

    "deck/deck"
)

func main(){

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
