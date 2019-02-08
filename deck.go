package deck

import (
	"math/rand"
	"sort"
)

// Suit defines a card's suit
type Suit int

const (
	spades Suit = iota
	diamonds
	clubs
	hearts
)

// Value defines a card's value
type Value int

const (
	ace Value = iota
	two
	three
	four
	five
	six
	seven
	eight
	nine
	jack
	queen
	king
)

var (
	suits  = []Suit{spades, diamonds, clubs, hearts}
	values = []Value{ace, two, three, four, five, six, seven, eight, nine, jack, queen, king}
)

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

// By is function predicate that compares card elements.
type By func(c1, c2 *Card) bool

// CardSorter holds a card deck and a list of comparison functions
// to sort the cards.
type CardSorter struct {
	deck Deck
	by   []By
}

// OrderedBy builds a CardSorter containing a list of comparison functions.
func OrderedBy(by ...By) *CardSorter {
	return &CardSorter{
		by: by,
	}
}

func bySuit(c1, c2 *Card) bool {
	return c1.Suit < c2.Suit
}
func byValue(c1, c2 *Card) bool {
	return c1.Value < c2.Value
}

// Sort sorts the deck of cards.
func (d *Deck) Sort() {
	OrderedBy(bySuit, byValue).sort(*d)
}

func (ms *CardSorter) sort(deck Deck) {
	ms.deck = deck
	sort.Sort(ms)
}

func (ms *CardSorter) Len() int {
	return len(ms.deck)
}

func (ms *CardSorter) Swap(i, j int) {
	ms.deck[i], ms.deck[j] = ms.deck[j], ms.deck[i]
}

func (ms *CardSorter) Less(i, j int) bool {
	p, q := &ms.deck[i], &ms.deck[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(ms.by)-1; k++ {
		compare := ms.by[k]
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
	return ms.by[k](p, q)
}
