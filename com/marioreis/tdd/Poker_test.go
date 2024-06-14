package tdd

import "testing"

func Test_CanInstantiateCard(t *testing.T) {
	card, err := NewCard(Spades, One)

	if err != nil || card.Value != 1 {
		t.Errorf("Incorrect Value!")
	}
}

func Test_AddCardToHand(t *testing.T) {
	card, err := NewCard(Spades, One)

	if err != nil || card.Value != 1 {
		t.Errorf("Incorrect Value!")
	}

	hand := new(Hand)
	hand.AddCard(card)

	if len(hand.cards) == 0 {
		t.Errorf("could not add card to hand")
	}
}

func Test_AddCardToHandAndOrder(t *testing.T) {
	card1, err := NewCard(Spades, One)

	if err != nil || card1.Value != 1 {
		t.Errorf("Incorrect Value!")
	}

	card2, err := NewCard(Spades, Two)

	if err != nil || card2.Value != 2 {
		t.Errorf("Incorrect Value!")
	}

	hand := new(Hand)
	hand.AddCard(card2)
	hand.AddCard(card1)

	if hand.cards[0].Value != 2 || hand.cards[1].Value != 1 {
		t.Errorf("add Card and order failed")
	}
}

func Test_AddCardToHandAndEvalHighestCard(t *testing.T) {
	card1, err := NewCard(Spades, One)

	if err != nil || card1.Value != 1 {
		t.Errorf("Incorrect Value!")
	}

	card2, err := NewCard(Spades, Two)

	if err != nil || card2.Value != 2 {
		t.Errorf("Incorrect Value!")
	}

	hand := new(Hand)
	hand.AddCard(card2)
	hand.AddCard(card1)

	if hand.highestCard != 2 {
		t.Errorf("Highest card failed")
	}
}

func Test_AddCardToHandAndEvalAllSuitsEqual(t *testing.T) {
	card1, err := NewCard(Spades, One)

	if err != nil || card1.Value != 1 {
		t.Errorf("Incorrect Value!")
	}

	card2, err := NewCard(Spades, Two)

	if err != nil || card2.Value != 2 {
		t.Errorf("Incorrect Value!")
	}

	hand := new(Hand)
	hand.AddCard(card2)
	hand.AddCard(card1)

	if hand.allSuitsEqual == false {
		t.Errorf("All suits equal faile!")
	}
}

func Test_AddCardToHandAndEvalAllSuitsNotEqual(t *testing.T) {
	card1, err := NewCard(Spades, One)

	if err != nil || card1.Value != 1 {
		t.Errorf("Incorrect Value!")
	}

	card2, err := NewCard(Hearts, Two)

	if err != nil || card2.Value != 2 {
		t.Errorf("Incorrect Value!")
	}

	hand := new(Hand)
	hand.AddCard(card2)
	hand.AddCard(card1)

	if hand.allSuitsEqual == true {
		t.Errorf("All suits equal faile!")
	}
}

func Test_AddCardToHandAndEvalIsSequence(t *testing.T) {
	card1, err := NewCard(Spades, One)

	if err != nil || card1.Value != 1 {
		t.Errorf("Incorrect Value!")
	}

	card2, err := NewCard(Hearts, Two)

	if err != nil || card2.Value != 2 {
		t.Errorf("Incorrect Value!")
	}

	card3, err := NewCard(Hearts, Three)

	if err != nil || card3.Value != 3 {
		t.Errorf("Incorrect Value!")
	}

	card4, err := NewCard(Hearts, Four)

	if err != nil || card4.Value != 4 {
		t.Errorf("Incorrect Value!")
	}

	card5, err := NewCard(Hearts, Five)

	if err != nil || card5.Value != 5 {
		t.Errorf("Incorrect Value!")
	}

	hand := new(Hand)
	hand.AddCard(card2)
	hand.AddCard(card1)
	hand.AddCard(card3)
	hand.AddCard(card4)
	hand.AddCard(card5)

	if hand.IsSequence() == false {
		t.Errorf("All suits equal faile!")
	}
}
