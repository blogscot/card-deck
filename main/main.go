package main

import (
	"fmt"
	"strings"

	"github.com/blogscot/deck"
)

func main() {
	cards := deck.New()
	cards.Shuffle()

	suit := func(c1, c2 *deck.Card) bool {
		return c1.Suit < c2.Suit
	}
	value := func(c1, c2 *deck.Card) bool {
		return c1.Value < c2.Value
	}

	deck.OrderedBy(suit, value).Sort(cards)

	for _, card := range cards {
		v := strings.Title((card.Value).String())
		s := strings.Title((card.Suit).String())
		fmt.Printf("%s of %s\n", v, s)
	}
}
