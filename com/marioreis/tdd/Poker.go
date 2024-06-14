package tdd

import (
	"fmt"
)

type Suit int
type Symbol string

const (
	Clubs Suit = iota
	Hearts
	Diamonds
	Spades
)

const MaxCardsCount = 5

const (
	One   Symbol = "1"
	Two   Symbol = "2"
	Three Symbol = "3"
	Four  Symbol = "4"
	Five  Symbol = "5"
	Six   Symbol = "6"
	Seven Symbol = "7"
	Eight Symbol = "8"
	Nine  Symbol = "9"
	Ten   Symbol = "10"
	Jack  Symbol = "J"
	Queen Symbol = "Q"
	King  Symbol = "K"
	Ace   Symbol = "A"
)

var cardValues = map[Symbol]int{
	One:   1,
	Two:   2,
	Three: 3,
	Four:  4,
	Five:  5,
	Six:   6,
	Seven: 7,
	Eight: 8,
	Nine:  9,
	Ten:   10,
	Jack:  11,
	Queen: 12,
	King:  13,
	Ace:   14,
}

type Card struct {
	Suit   Suit
	Symbol Symbol
	Value  int
}

func NewCard(suit Suit, symbol Symbol) (*Card, error) {
	value, found := cardValues[symbol]

	if !found {
		return nil, fmt.Errorf("symbol %s not found", symbol)
	}

	return &Card{
		Symbol: symbol,
		Suit:   suit,
		Value:  value,
	}, nil
}

type Hand struct {
	cards         []*Card
	highestCard   int
	allSuitsEqual bool
	suits         map[Suit]int
}

func (h *Hand) AddCard(card *Card) {
	h.cards = append(h.cards, card)

	for i := len(h.cards) - 1; i >= 1; i-- {
		if h.cards[i].Value > h.cards[i-1].Value {
			swap(h.cards, i, i-1)
		}
	}

	if card.Value > h.highestCard {
		h.highestCard = card.Value
	}

	if h.suits == nil {
		h.suits = make(map[Suit]int)
	}

	suitCount, exists := h.suits[card.Suit]

	if exists {
		h.suits[card.Suit] = suitCount + 1
	} else {
		h.suits[card.Suit] = 1
	}

	h.allSuitsEqual = len(h.suits) == 1
}

func (h *Hand) IsSequence() bool {
	if len(h.cards) < 5 {
		return false
	}

	expectedValue := h.highestCard

	for i := range h.cards {
		if expectedValue != h.cards[i].Value {
			return false
		}
		expectedValue--
	}

	return true
}

func swap(array []*Card, higherIndex int, lowerIndex int) {
	temp := array[lowerIndex]
	array[lowerIndex] = array[higherIndex]
	array[higherIndex] = temp
}
