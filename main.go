package main

import "fmt"

func main() {
	// sets cards equal to a slice of strings
	// cards := []string{newCard(), newCard()}

	// Referencing the deck we created in deck.go. Allows us to add receiver functions
	// to allow for more methods
	cards := deck{newCard(), newCard()}
	cards = append(cards, "Six of Spades")


	for idx, card := range cards {
		fmt.Println(idx, card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}