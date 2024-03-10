package deck

import (
	"testing"
)

func compareDecks(d1, d2 []Card) bool {
	if len(d1) != len(d2) {
		return false
	}
	for i := range len(d1) {
		c1, c2 := d1[i], d2[i]
		if c1.Suit != c2.Suit && c1.Value != c2.Value {
			return false
		}
	}
	return true
}

func TestNew(t *testing.T) {
	cards := NewDeck()
	if len(cards) != 52 {
		t.Error("Wrong number of cards in a new deck.")
	}
}

func TestString(t *testing.T) {
	card := Card{Suit: Hearts, Value: Five}
	expected := "Five of Hearts"

	if card.String() != expected {
		t.Errorf("Strings dont match - %s != %s", card.String(), expected)
	}

	card = Card{Suit: Joker, Value: 0}
	expected = "Joker"

	if card.String() != expected {
		t.Errorf("Strings dont match - %s != %s", card.String(), expected)
	}
}

func TestSortSimple(t *testing.T) {
	deck := []Card{
		Card{Hearts, Five},
		Card{Hearts, Four},
		Card{Spades, Ace},
		Card{Clubs, King},
	}
	expDeck := []Card{
		Card{Spades, Ace},
		Card{Clubs, King},
		Card{Hearts, Four},
		Card{Hearts, Five},
	}
	SortDeckSimple(deck)
	if !compareDecks(deck, expDeck) {
		t.Error("SortedDeck is not the same as expected")
	}
}

func TestAddJoker(t *testing.T) {
	deck := NewDeck()
	deck = AddJoker(deck)

	if deck[len(deck)-1].Suit != Joker {
		t.Error("Expected last Card in deck to be a Joker")
	}

	if len(deck) != 53 {
		t.Error("Wrong number of cards in a new deck with 1 Joker")
	}

}
