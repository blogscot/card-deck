package main

import (
	"fmt"
	"strings"

	"github.com/blogscot/deck"
)

func main() {
	cards := deck.New()
	cards.Shuffle()
	cards.Sort()

	for _, card := range cards {
		v := strings.Title((card.Value).String())
		s := strings.Title((card.Suit).String())
		fmt.Printf("%s of %s\n", v, s)
	}
}
