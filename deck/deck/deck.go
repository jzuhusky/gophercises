//go:generate stringer -type=SuitType,ValueType -output=type_strings.go
package deck

import (
    "fmt"
	"math/rand"
	"sort"
	"time"
)

type SuitType int
type ValueType int

const (
	// Suits
	// Given numerical value for reasons of sorting
	// Any game-specific logic should not depend on the int values specified here
	// But rather be a function of the constant name
	Spades   SuitType = 1
	Diamonds SuitType = 2
	Clubs    SuitType = 3
	Hearts   SuitType = 4
	Joker    SuitType = 5

	// Values
	// Given numerical value for reasons of sorting
	// Any game-specific logic should not depend on the int values specified here
	// But rather be a function of the constant name
	Ace   ValueType = 1
	Two   ValueType = 2
	Three ValueType = 3
	Four  ValueType = 4
	Five  ValueType = 5
	Six   ValueType = 6
	Seven ValueType = 7
	Eight ValueType = 8
	Nine  ValueType = 9
	Ten   ValueType = 10
	Jack  ValueType = 11
	Queen ValueType = 12
	King  ValueType = 13
)

type Card struct {
	Suit  SuitType
	Value ValueType
}

func (c Card) String() string {
    if c.Suit == Joker {
        return c.Suit.String()
    }
    return fmt.Sprintf("%s of %s", c.Value.String(), c.Suit.String())
}

func NewDeck() []Card {
	suits := []SuitType{Spades, Hearts, Clubs, Diamonds}
	values := []ValueType{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
	cards := make([]Card, 52)
	idx := 0
	for _, suit := range suits {
		for _, value := range values {
			cards[idx] = Card{suit, value}
			idx += 1
		}
	}
	return cards
}

func SortDeckSimple(deck []Card) {
	sort.Slice(deck, func(i, j int) bool {
		c1, c2 := deck[i], deck[j]
		if c1.Suit == c2.Suit {
			return c1.Value < c2.Value
		}
		if c1.Suit < c2.Suit {
			return true
		}
		return false
	})
}

func AddJoker(deck []Card) []Card {
    return append(deck, Card{Suit: Joker, Value: 0})
}

func ShuffleDeck(deck []Card) []Card {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
	return deck
}
