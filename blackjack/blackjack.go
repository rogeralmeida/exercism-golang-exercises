package blackjack

const blackJack = 21
const hit = "H"
const stand = "S"
const split = "P"
const win = "W"

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cards := map[string]int{
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
		"ace":   11,
	}
	cardValue := cards[card]
	return cardValue
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	action := hit
	totalMyCards := ParseCard(card1) + ParseCard(card2)
	totalDealerCard := ParseCard(dealerCard)
	switch {
	case totalMyCards > blackJack:
		action = split
	case totalMyCards == blackJack && totalDealerCard < 10:
		action = win
	case totalMyCards == blackJack && totalDealerCard >= 10:
		action = stand
	case totalMyCards >= 17 && totalMyCards <= 20:
		action = stand
	case totalMyCards >= 12 && totalMyCards <= 16 && totalDealerCard < 7:
		action = stand
	case totalMyCards >= 12 && totalMyCards <= 16 && totalDealerCard >= 7:
		action = hit
	}
	return action
}
