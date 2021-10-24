package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type of "dec"
//which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	//for Each suit element add of and value element
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

//method to the type deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//method to make slices of cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}

//Helper function which takes deck and returns string representation of it
func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

//Saving data to the Hard Drive
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		//exiting program with another go package
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)

}

//Shuffle function that uses random numbers to change their index in the oreder to change theri position

func (d deck) shuffle() {
	//creating seed value, value inside is UnixNano starting tamo which will give us different time/number every time
	source := rand.NewSource(time.Now().UnixNano())

	//creating new rand object from seed value which is of type rand and can be use as rand(default package object)
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]

	}

}
