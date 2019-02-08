package main

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
	suit  Suit
	value Value
}

// Deck defines a deck of cards
type Deck []Card
