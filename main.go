package main

func main() {
	// sets cards equal to a slice of strings
	// cards := []string{newCard(), newCard()}

	// Referencing the deck we created in deck.go. Allows us to add receiver functions
	// to allow for more methods
	cards := newDeck()

	// specify both values that deal() will return
	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
	// References receiver function created in deck.go to loop & print
	// cards.print()
}
