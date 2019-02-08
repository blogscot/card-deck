package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main() {
	deck := New()
	deck.Shuffle()
	// for _, card := range deck {
	// 	v := strings.Title((card.value).String())
	// 	s := strings.Title((card.suit).String())
	// 	fmt.Printf("%s of %s\n", v, s)
	// }
	// deck.BySuit()
	// sort.Sort(BySuit(deck))

	suit := func(c1, c2 *Card) bool {
		return c1.suit < c2.suit
	}
	value := func(c1, c2 *Card) bool {
		return c1.value < c2.value
	}

	OrderedBy(suit, value).Sort(deck)

	for _, card := range deck {
		v := strings.Title((card.value).String())
		s := strings.Title((card.suit).String())
		fmt.Printf("%s of %s\n", v, s)
	}
}

// New generates a deck of cards
func New() (deck Deck) {
	for _, s := range suits {
		for _, v := range values {
			deck = append(deck, Card{suit: s, value: v})
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

// Sort performs sorting using the CardSorter.
func (ms *CardSorter) Sort(deck Deck) {
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
