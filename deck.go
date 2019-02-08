// Package deck provides a means to create, shuffle, sort, and show a deck of cards.
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Suit defines a card's suit
type Suit int

const (
	Spades Suit = iota
	Diamonds
	Clubs
	Hearts
)

// Value defines a card's value
type Value int

const (
	Ace Value = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

var (
	suits  = []Suit{Spades, Diamonds, Clubs, Hearts}
	values = []Value{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Card is a card in a deck of cards
type Card struct {
	Suit  Suit
	Value Value
}

// Deck defines a deck of cards
type Deck []Card

// New generates a deck of cards
func New() (deck Deck) {
	for _, s := range suits {
		for _, v := range values {
			deck = append(deck, Card{Suit: s, Value: v})
		}
	}
	return
}

// Shuffle randomises the order of a deck of cards.
func (d *Deck) Shuffle() {
	for i := range *d {
		j := rand.Intn(i + 1)
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	}
	return
}

// Sort sorts the deck of cards.
func (d *Deck) Sort() {
	orderedBy(bySuit, byValue).sort(*d)
}

// Show prints the deck of cards.
func Show(cards []Card) {
	for _, card := range cards {
		v := card.Value
		s := card.Suit
		fmt.Printf("%s of %s\n", v, s)
	}
}

// compare is function predicate that compares card elements.
type compare func(c1, c2 *Card) bool

// cardSorter holds a card deck and a list of comparison functions
// to sort the cards.
type cardSorter struct {
	deck    Deck
	compare []compare
}

// orderedBy builds a CardSorter containing a list of comparison functions.
func orderedBy(by ...compare) *cardSorter {
	return &cardSorter{
		compare: by,
	}
}

func bySuit(c1, c2 *Card) bool {
	return c1.Suit < c2.Suit
}
func byValue(c1, c2 *Card) bool {
	return c1.Value < c2.Value
}

func (ms *cardSorter) sort(deck Deck) {
	ms.deck = deck
	sort.Sort(ms)
}

func (ms *cardSorter) Len() int {
	return len(ms.deck)
}

func (ms *cardSorter) Swap(i, j int) {
	ms.deck[i], ms.deck[j] = ms.deck[j], ms.deck[i]
}

func (ms *cardSorter) Less(i, j int) bool {
	p, q := &ms.deck[i], &ms.deck[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(ms.compare)-1; k++ {
		compare := ms.compare[k]
		switch {
		case compare(p, q):
			// p < q, so we have a decision.
			return true
		case compare(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.compare[k](p, q)
}
