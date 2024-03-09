//go:generate echo "hello world"
package deck

type Suit int

const (
    Spades Suit = iota
    Hearts Suit = iota
    Clubs Suit = iota
    Diomonds Suit= iota
)

type Card struct {
    Value int
    Suit 
}
