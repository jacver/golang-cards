package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

// Turning type deck into String, add receiver so we can eventually call cards.toString()
func (d deck) toString() string {
	// type conversion, we want a slice of strings and we have d
	return strings.Join([]string(d), ",")
	// returns our string slice as a single string seperated by commas 
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// When handling a Go error, consider what you want to happen when something breaks
		// Option 1: Log the error and return a call to newDeck() so user can still have a deck
		// Option 2: Log error and quit program
		fmt.Println("Error: ", err)
		// Any int except 0 will trigger an exit -- 1 is fine here
		os.Exit(1)
	} 
	
	// turn byte slice into slice of strings
	s := strings.Split(string(bs), ",") 
	// Type conversion to get a deck from our slice of strings
	return deck(s)
}

// ideally we could cards.shuffle() -- so we need receiver
func (d deck) shuffle() {
	// generating a truly random number
	source := rand.NewSource(time.Now().UnixNano()) // each time we run the program, it will generate a seed bc time will always be different
	r := rand.New(source)
	// r is type Rand so now it can call Intn below

	for i := range d {
		// no need for card, just use index
		newPosition := r.Intn(len(d) - 1) 


		d[i], d[newPosition] = d[newPosition], d[i] // swaps both elements
	}
}