package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	// create and return a list of playing cards (a deck)
	// REMEMBER: must explicitly state type of return value (deck)
	cards := deck{}

	// The deck will eventually have all 52 cards. Best way to approach this is two slices
	// slice 1: contains all suits  ---- slice 2: contains all values
	// with a nested for loop we can combine suit & value

	cardSuits := []string{"Spades", "diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func (d deck) print() {
	// (d deck) is the receiver for this function - name as 1st letter of type
	// d is a reference actual copy of the deck we're using (like <this> in javascript)
	// // in cards.print() imagine that cards replaces the d

	//  SUMMARY:
	// // (d deck) is of type deck
	// // this method is now available for anything of type deck
	// // we can run cards.print() because it is equal to type deck
	for idx, card := range d {
		fmt.Println(idx, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	// set to return TWO values, both of type deck
	// splitting the deck into two, one of handsize, one of everything else
	return d[:handSize], d[handSize:]
}