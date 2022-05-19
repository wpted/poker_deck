package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

// Create a new type of deck
// which is a slice of strings
type deck []string // let the deck act the same as the slice of strings

// Create and return a list of playing cards, Essentially an array of strings.
// Don't need a receiver since creating deck is what you do right now
// have nothing to receive
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// log out the contents of a deck of cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Shuffles all the cards in a deck
// For each index(card in cards), generate a random number between 0 and len(cards) - 1
// Swap the current card and the card at cards[random_number]
func (d deck) shuffle() {
	for i := range d {
		// Seeding with the same value results in the same random sequence each run.
		// For different numbers, seed with a different value, such as
		// time.Now().UnixNano(), which yields a constantly-changing number(int64).
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		// Random integer within the whole deck
		newPosition := r.Intn(len(d))
		// Swapping
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// Create a hand of cards
// Slicing [0:handSize] makes it work
func deal(d deck, handSize int) (deck, deck) {
	return d[0:handSize], d[handSize:]

}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Save a list of cards to a file on the local machine
// func WriteFile(filename string, data []byte, perm fs.FileMode) error
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Reads the file from the stored byte slice
func newDeckFromFile(filename string) deck {
	// Returns a byte string and an err
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		// basically the same as the below
		log.Fatal("Error: ", err)
		//fmt.Println("Error", err)
		//os.Exit(1)
	}
	return deck(strings.Split(string(byteSlice), ","))
}
