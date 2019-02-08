package main

import (
	"github.com/blogscot/deck"
)

func main() {
	cards := deck.New()
	cards.Shuffle()
	cards.Sort()
	deck.Show(cards)
}
