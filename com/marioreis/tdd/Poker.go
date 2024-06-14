package tdd

import (
	"fmt"
)

type Suit int
type Symbol string
type WinnerHand int

const (
	Clubs Suit = iota
	Hearts
	Diamonds
	Spades
)

const MaxCardsCount = 5

const (
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

const (
	RoyalFlush    WinnerHand = 1000
	StraightFlush WinnerHand = 900
	FourOfAKind   WinnerHand = 800
	FullHouse     WinnerHand = 700
	Flush         WinnerHand = 600
	ThreeOfAKind  WinnerHand = 500
	TwoPair       WinnerHand = 400
	Pair          WinnerHand = 300
	HighCard      WinnerHand = 200
)

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
	sameOfAKind   map[Symbol]int
	winnerHand    WinnerHand
}

func order(h *Hand) {
	for i := len(h.cards) - 1; i >= 1; i-- {
		if h.cards[i].Value > h.cards[i-1].Value {
			swap(h.cards, i, i-1)
		}
	}

	if isUnorderedSequenceWhereAceValuesOne(h.cards) {
		ace := h.cards[0]
		h.cards = append(h.cards, ace)
		h.cards = h.cards[1:]
	}
}

func (h *Hand) AddCard(card *Card) {
	h.cards = append(h.cards, card)

	if len(h.cards) >= 2 {
		order(h)
	}

	if card.Value > h.highestCard {
		h.highestCard = card.Value
	}

	if h.suits == nil {
		h.suits = make(map[Suit]int)
	}

	if h.sameOfAKind == nil {
		h.sameOfAKind = make(map[Symbol]int)
	}

	suitCount, exists := h.suits[card.Suit]

	if exists {
		h.suits[card.Suit] = suitCount + 1
	} else {
		h.suits[card.Suit] = 1
	}

	sameOfAKindCount, sameExists := h.sameOfAKind[card.Symbol]

	if sameExists {
		h.sameOfAKind[card.Symbol] = sameOfAKindCount + 1
	} else {
		h.sameOfAKind[card.Symbol] = 1
	}

	h.allSuitsEqual = len(h.suits) == 1

	if isRoyalFlush(h) {
		h.winnerHand = RoyalFlush
	}
}

func isRoyalFlush(h *Hand) bool {
	return h.allSuitsEqual && h.IsSequence() && h.cards[0].Symbol == "A"
}

func isStraightFlush(h *Hand) bool {
	return h.allSuitsEqual && h.IsSequence() && h.cards[0].Symbol != "A"
}

func isFullHouse(hand *Hand) bool {
	threeOfAKind := false
	twoOfAKind := false

	for i := range hand.sameOfAKind {
		if hand.sameOfAKind[i] == 3 {
			threeOfAKind = true
		}

		if hand.sameOfAKind[i] == 2 {
			twoOfAKind = true
		}
	}

	return threeOfAKind && twoOfAKind
}

func isFlush(hand *Hand) bool {
	return hand.allSuitsEqual && !hand.IsSequence()
}

func isStraight(hand *Hand) bool {
	return hand.IsSequence() && !hand.allSuitsEqual
}

func isThreeOfAKind(hand *Hand) bool {
	threeOfAKind := false
	twoOfAKind := false

	for i := range hand.sameOfAKind {
		if hand.sameOfAKind[i] == 3 {
			threeOfAKind = true
		}

		if hand.sameOfAKind[i] == 2 {
			twoOfAKind = true
		}
	}

	return threeOfAKind && !twoOfAKind
}

func isUnorderedSequenceWhereAceValuesOne(cards []*Card) bool {
	return cards[0].Value == 14 && cards[1].Value == 5 && cards[2].Value == 4 && cards[3].Value == 3 &&
		cards[4].Value == 2
}

func isOrderedSequenceWhereAceValuesOne(cards []*Card) bool {
	return cards[0].Value == 5 && cards[1].Value == 4 && cards[2].Value == 3 &&
		cards[3].Value == 2 && cards[4].Value == 14
}

func isFourOfAKind(hand *Hand) bool {
	for i := range hand.sameOfAKind {
		if hand.sameOfAKind[i] == 4 {
			return true
		}
	}

	return false
}

func isTwoPair(hand *Hand) bool {
	symbols := make(map[Symbol]int)

	for i := range hand.sameOfAKind {
		if hand.sameOfAKind[i] == 2 {
			symbols[i] = 1
		}
	}

	return len(symbols) == 2
}

func isPair(hand *Hand) bool {
	if isFullHouse(hand) {
		return false
	}

	symbols := make(map[Symbol]int)

	for i := range hand.sameOfAKind {
		if hand.sameOfAKind[i] == 2 {
			symbols[i] = 1
		}
	}

	return len(symbols) == 1
}

func isHighestCard(hand *Hand) bool {
	isSequence := hand.IsSequence()
	allSuitsEqual := hand.allSuitsEqual
	pair := isPair(hand)
	twoPair := isTwoPair(hand)
	fullHouse := isFullHouse(hand)
	threeOfAKind := isThreeOfAKind(hand)
	fourOfAKind := isFourOfAKind(hand)

	return !isSequence && !allSuitsEqual && !pair && !twoPair && !fullHouse && !threeOfAKind &&
		!fourOfAKind
}

func (h *Hand) IsSequence() bool {
	if len(h.cards) < 5 {
		return false
	}

	expectedValue := h.highestCard

	if isOrderedSequenceWhereAceValuesOne(h.cards) {
		return true
	}

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
