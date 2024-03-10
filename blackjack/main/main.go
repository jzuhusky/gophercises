package main 

import (
    "fmt"

    "github.com/jzuhusky/gophercises/deck"
)

func main(){
    fmt.Println("hello world")
    deck := NewDeck()
    fmt.Println(len(deck))
}
