package tdd

import "testing"

func Test_CanInstantiateCard(t *testing.T) {
	card, err := NewCard(Spades, Two)

	if err != nil || card.Value != 2 {
		t.Errorf("Incorrect Value!")
	}
}

func Test_AddCardToHand(t *testing.T) {
	card, err := NewCard(Spades, Two)

	if err != nil || card.Value != 2 {
		t.Errorf("Incorrect Value!")
	}

	hand := new(Hand)
	hand.AddCard(card)

	if len(hand.cards) == 0 {
		t.Errorf("could not add card to hand")
	}
}

func Test_AddCardToHandAndOrder(t *testing.T) {
	card1, err := NewCard(Spades, Two)

	if err != nil || card1.Value != 2 {
		t.Errorf("Incorrect Value!")
	}

	card2, err := NewCard(Spades, Three)

	if err != nil || card2.Value != 3 {
		t.Errorf("Incorrect Value!")
	}

	hand := new(Hand)
	hand.AddCard(card2)
	hand.AddCard(card1)

	if hand.cards[0].Value != 3 || hand.cards[1].Value != 2 {
		t.Errorf("add Card and order failed")
	}
}

func Test_AddCardToHandAndEvalHighestCard(t *testing.T) {
	card1, err := NewCard(Spades, Two)

	if err != nil || card1.Value != 2 {
		t.Errorf("Incorrect Value!")
	}

	card2, err := NewCard(Spades, Three)

	if err != nil || card2.Value != 3 {
		t.Errorf("Incorrect Value!")
	}

	hand := new(Hand)
	hand.AddCard(card2)
	hand.AddCard(card1)

	if hand.highestCard != 3 {
		t.Errorf("Highest card failed")
	}
}

func Test_AddCardToHandAndEvalAllSuitsEqual(t *testing.T) {
	card1, err := NewCard(Spades, Two)

	if err != nil || card1.Value != 2 {
		t.Errorf("Incorrect Value!")
	}

	card2, err := NewCard(Spades, Three)

	if err != nil || card2.Value != 3 {
		t.Errorf("Incorrect Value!")
	}

	hand := new(Hand)
	hand.AddCard(card2)
	hand.AddCard(card1)

	if hand.allSuitsEqual == false {
		t.Errorf("All suits equal failed!")
	}
}

func Test_AddCardToHandAndEvalAllSuitsNotEqual(t *testing.T) {
	card1, err := NewCard(Spades, Two)

	if err != nil || card1.Value != 2 {
		t.Errorf("Incorrect Value!")
	}

	card2, err := NewCard(Hearts, Three)

	if err != nil || card2.Value != 3 {
		t.Errorf("Incorrect Value!")
	}

	hand := new(Hand)
	hand.AddCard(card2)
	hand.AddCard(card1)

	if hand.allSuitsEqual == true {
		t.Errorf("All suits equal failed!")
	}
}

func Test_AddCardToHandAndEvalIsSequence(t *testing.T) {
	card1, err := NewCard(Spades, Two)

	if err != nil || card1.Value != 2 {
		t.Errorf("Incorrect Value!")
	}

	card2, err := NewCard(Hearts, Three)

	if err != nil || card2.Value != 3 {
		t.Errorf("Incorrect Value!")
	}

	card3, err := NewCard(Hearts, Four)

	if err != nil || card3.Value != 4 {
		t.Errorf("Incorrect Value!")
	}

	card4, err := NewCard(Hearts, Five)

	if err != nil || card4.Value != 5 {
		t.Errorf("Incorrect Value!")
	}

	card5, err := NewCard(Hearts, Six)

	if err != nil || card5.Value != 6 {
		t.Errorf("Incorrect Value!")
	}

	hand := new(Hand)
	hand.AddCard(card2)
	hand.AddCard(card1)
	hand.AddCard(card3)
	hand.AddCard(card4)
	hand.AddCard(card5)

	if hand.IsSequence() == false {
		t.Errorf("Is sequence failed!")
	}
}

func Test_AddCardToHandAndEvalIsUnorderedSequenceWithAceValuesOne(t *testing.T) {
	card1, err := NewCard(Spades, Two)

	if err != nil || card1.Value != 2 {
		t.Errorf("Incorrect Value!")
	}

	card2, err := NewCard(Hearts, Three)

	if err != nil || card2.Value != 3 {
		t.Errorf("Incorrect Value!")
	}

	card3, err := NewCard(Hearts, Four)

	if err != nil || card3.Value != 4 {
		t.Errorf("Incorrect Value!")
	}

	card4, err := NewCard(Hearts, Five)

	if err != nil || card4.Value != 5 {
		t.Errorf("Incorrect Value!")
	}

	card5, err := NewCard(Hearts, Ace)

	if err != nil || card5.Value != 14 {
		t.Errorf("Incorrect Value!")
	}

	hand := new(Hand)
	hand.AddCard(card2)
	hand.AddCard(card1)
	hand.AddCard(card3)
	hand.AddCard(card4)
	hand.AddCard(card5)

	if hand.IsSequence() == false {
		t.Errorf("Is sequence failed!")
	}
}

func Test_IsHandRoyalFlush(t *testing.T) {
	card1, err := NewCard(Hearts, Ace)

	if err != nil || card1.Value != 14 {
		t.Errorf("Incorrect Value!")
	}

	card2, err := NewCard(Hearts, King)

	if err != nil || card2.Value != 13 {
		t.Errorf("Incorrect Value!")
	}

	card3, err := NewCard(Hearts, Queen)

	if err != nil || card3.Value != 12 {
		t.Errorf("Incorrect Value!")
	}

	card4, err := NewCard(Hearts, Jack)

	if err != nil || card4.Value != 11 {
		t.Errorf("Incorrect Value!")
	}

	card5, err := NewCard(Hearts, Ten)

	if err != nil || card5.Value != 10 {
		t.Errorf("Incorrect Value!")
	}

	hand := new(Hand)
	hand.AddCard(card2)
	hand.AddCard(card1)
	hand.AddCard(card3)
	hand.AddCard(card4)
	hand.AddCard(card5)

	if hand.IsSequence() == false && !isRoyalFlush(hand) {
		t.Errorf("Is royal flush failed!")
	}

}

func Test_order(t *testing.T) {

	card1, err := NewCard(Hearts, Ace)

	if err != nil || card1.Value != 14 {
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

	order(hand)

	if hand.cards[0].Value != 5 || hand.cards[1].Value != 4 || hand.cards[2].Value != 3 ||
		hand.cards[3].Value != 2 || hand.cards[4].Value != 14 {
		t.Errorf("order cards failed")
	}

}

func Test_Straight_Flush(t *testing.T) {

	card1, err := NewCard(Hearts, Ace)

	if err != nil || card1.Value != 14 {
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
	hand.AddCard(card5)
	hand.AddCard(card4)

	if !isStraightFlush(hand) {
		t.Errorf("straight flush failed")
	}

}

func Test_FourOfAKind(t *testing.T) {

	card1, err := NewCard(Hearts, Ace)

	if err != nil || card1.Value != 14 {
		t.Errorf("Incorrect Value!")
	}

	card2, err := NewCard(Clubs, Ace)

	if err != nil || card2.Value != 14 {
		t.Errorf("Incorrect Value!")
	}

	card3, err := NewCard(Spades, Ace)

	if err != nil || card3.Value != 14 {
		t.Errorf("Incorrect Value!")
	}

	card4, err := NewCard(Diamonds, Ace)

	if err != nil || card4.Value != 14 {
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
	hand.AddCard(card5)
	hand.AddCard(card4)

	if !isFourOfAKind(hand) {
		t.Errorf("straight flush failed")
	}

}

func Test_FullHouse(t *testing.T) {
	card1, err := NewCard(Hearts, Ace)

	if err != nil || card1.Value != 14 {
		t.Errorf("Incorrect Value!")
	}

	card2, err := NewCard(Clubs, Ace)

	if err != nil || card2.Value != 14 {
		t.Errorf("Incorrect Value!")
	}

	card3, err := NewCard(Spades, Ace)

	if err != nil || card3.Value != 14 {
		t.Errorf("Incorrect Value!")
	}

	card4, err := NewCard(Diamonds, King)

	if err != nil || card4.Value != 13 {
		t.Errorf("Incorrect Value!")
	}

	card5, err := NewCard(Hearts, King)

	if err != nil || card5.Value != 13 {
		t.Errorf("Incorrect Value!")
	}

	hand := new(Hand)
	hand.AddCard(card2)
	hand.AddCard(card1)
	hand.AddCard(card3)
	hand.AddCard(card5)
	hand.AddCard(card4)

	if !isFullHouse(hand) {
		t.Errorf("Full House failed")
	}

}

func Test_Flush(t *testing.T) {
	card1, err := NewCard(Hearts, Ace)

	if err != nil || card1.Value != 14 {
		t.Errorf("Incorrect Value!")
	}

	card2, err := NewCard(Hearts, Seven)

	if err != nil || card2.Value != 7 {
		t.Errorf("Incorrect Value!")
	}

	card3, err := NewCard(Hearts, Five)

	if err != nil || card3.Value != 5 {
		t.Errorf("Incorrect Value!")
	}

	card4, err := NewCard(Hearts, Three)

	if err != nil || card4.Value != 3 {
		t.Errorf("Incorrect Value!")
	}

	card5, err := NewCard(Hearts, Two)

	if err != nil || card5.Value != 2 {
		t.Errorf("Incorrect Value!")
	}

	hand := new(Hand)
	hand.AddCard(card2)
	hand.AddCard(card1)
	hand.AddCard(card3)
	hand.AddCard(card5)
	hand.AddCard(card4)

	if !isFlush(hand) {
		t.Errorf("Flush failed")
	}
}

func Test_Straight(t *testing.T) {
	card1, err := NewCard(Hearts, Six)

	if err != nil || card1.Value != 6 {
		t.Errorf("Incorrect Value!")
	}

	card2, err := NewCard(Clubs, Seven)

	if err != nil || card2.Value != 7 {
		t.Errorf("Incorrect Value!")
	}

	card3, err := NewCard(Spades, Eight)

	if err != nil || card3.Value != 8 {
		t.Errorf("Incorrect Value!")
	}

	card4, err := NewCard(Diamonds, Nine)

	if err != nil || card4.Value != 9 {
		t.Errorf("Incorrect Value!")
	}

	card5, err := NewCard(Hearts, Ten)

	if err != nil || card5.Value != 10 {
		t.Errorf("Incorrect Value!")
	}

	hand := new(Hand)
	hand.AddCard(card2)
	hand.AddCard(card1)
	hand.AddCard(card3)
	hand.AddCard(card5)
	hand.AddCard(card4)

	if !isStraight(hand) {
		t.Errorf("Straight failed")
	}
}

func Test_ThreeOfAKind(t *testing.T) {
	card1, err := NewCard(Hearts, Ace)

	if err != nil || card1.Value != 14 {
		t.Errorf("Incorrect Value!")
	}

	card2, err := NewCard(Clubs, Ace)

	if err != nil || card2.Value != 14 {
		t.Errorf("Incorrect Value!")
	}

	card3, err := NewCard(Spades, Ace)

	if err != nil || card3.Value != 14 {
		t.Errorf("Incorrect Value!")
	}

	card4, err := NewCard(Diamonds, Nine)

	if err != nil || card4.Value != 9 {
		t.Errorf("Incorrect Value!")
	}

	card5, err := NewCard(Hearts, Ten)

	if err != nil || card5.Value != 10 {
		t.Errorf("Incorrect Value!")
	}

	hand := new(Hand)
	hand.AddCard(card2)
	hand.AddCard(card1)
	hand.AddCard(card3)
	hand.AddCard(card5)
	hand.AddCard(card4)

	if !isThreeOfAKind(hand) {
		t.Errorf("isThreeOfAKind failed")
	}
}

func Test_TwoPair(t *testing.T) {
	card1, err := NewCard(Hearts, Ace)

	if err != nil || card1.Value != 14 {
		t.Errorf("Incorrect Value!")
	}

	card2, err := NewCard(Clubs, Ace)

	if err != nil || card2.Value != 14 {
		t.Errorf("Incorrect Value!")
	}

	card3, err := NewCard(Spades, King)

	if err != nil || card3.Value != 13 {
		t.Errorf("Incorrect Value!")
	}

	card4, err := NewCard(Diamonds, King)

	if err != nil || card4.Value != 13 {
		t.Errorf("Incorrect Value!")
	}

	card5, err := NewCard(Hearts, Ten)

	if err != nil || card5.Value != 10 {
		t.Errorf("Incorrect Value!")
	}

	hand := new(Hand)
	hand.AddCard(card2)
	hand.AddCard(card1)
	hand.AddCard(card3)
	hand.AddCard(card5)
	hand.AddCard(card4)

	if !isTwoPair(hand) {
		t.Errorf("isTwoPair failed")
	}
}

func Test_Pair(t *testing.T) {
	card1, err := NewCard(Hearts, Ace)

	if err != nil || card1.Value != 14 {
		t.Errorf("Incorrect Value!")
	}

	card2, err := NewCard(Clubs, Ace)

	if err != nil || card2.Value != 14 {
		t.Errorf("Incorrect Value!")
	}

	card3, err := NewCard(Spades, King)

	if err != nil || card3.Value != 13 {
		t.Errorf("Incorrect Value!")
	}

	card4, err := NewCard(Diamonds, Jack)

	if err != nil || card4.Value != 11 {
		t.Errorf("Incorrect Value!")
	}

	card5, err := NewCard(Hearts, Ten)

	if err != nil || card5.Value != 10 {
		t.Errorf("Incorrect Value!")
	}

	hand := new(Hand)
	hand.AddCard(card2)
	hand.AddCard(card1)
	hand.AddCard(card3)
	hand.AddCard(card5)
	hand.AddCard(card4)

	if !isPair(hand) {
		t.Errorf("isPair failed")
	}
}

func Test_NotPair(t *testing.T) {
	card1, err := NewCard(Hearts, Ace)

	if err != nil || card1.Value != 14 {
		t.Errorf("Incorrect Value!")
	}

	card2, err := NewCard(Clubs, Ace)

	if err != nil || card2.Value != 14 {
		t.Errorf("Incorrect Value!")
	}

	card3, err := NewCard(Spades, King)

	if err != nil || card3.Value != 13 {
		t.Errorf("Incorrect Value!")
	}

	card4, err := NewCard(Diamonds, King)

	if err != nil || card4.Value != 13 {
		t.Errorf("Incorrect Value!")
	}

	card5, err := NewCard(Hearts, King)

	if err != nil || card5.Value != 13 {
		t.Errorf("Incorrect Value!")
	}

	hand := new(Hand)
	hand.AddCard(card2)
	hand.AddCard(card1)
	hand.AddCard(card3)
	hand.AddCard(card5)
	hand.AddCard(card4)

	if isPair(hand) {
		t.Errorf("isNotPair failed")
	}
}

func Test_HighCard(t *testing.T) {
	card1, err := NewCard(Hearts, Ace)

	if err != nil || card1.Value != 14 {
		t.Errorf("Incorrect Value!")
	}

	card2, err := NewCard(Clubs, King)

	if err != nil || card2.Value != 13 {
		t.Errorf("Incorrect Value!")
	}

	card3, err := NewCard(Spades, Jack)

	if err != nil || card3.Value != 11 {
		t.Errorf("Incorrect Value!")
	}

	card4, err := NewCard(Diamonds, Queen)

	if err != nil || card4.Value != 12 {
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
	hand.AddCard(card5)
	hand.AddCard(card4)

	if !isHighestCard(hand) {
		t.Errorf("isHighestCard failed")
	}
}
