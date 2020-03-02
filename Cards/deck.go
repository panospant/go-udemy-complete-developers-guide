package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// instead of naming the indexes of the loops we added _. This means something like:
	//  we know we have a value there, we just don't care about the value
	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}

	return cards
}

// d deck part is called a receiver
// this means any variable of type "deck" can have access to this print method
func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

// this functions returns two values of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) toByteSlice() []byte {
	return []byte(d.toString())
}

func (d deck) saveToFile(filename string) error {

	// the last part is the permission of the file and it means anyone can read write this file etc.
	return ioutil.WriteFile(filename, d.toByteSlice(), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)

	if err != nil {
		// log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	var s = strings.Split(string(byteSlice), ",")

	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	var r = rand.New(source)
	for index := range d {
		newPosition := r.Intn(len(d) - 1)

		// this is a way to swap positions in a slice
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
