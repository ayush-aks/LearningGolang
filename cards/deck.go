package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new type of 'deck' which is a slice of strings.
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Diamond", "Heart", "Spades", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, value := range cardValues {
		for _, suits := range cardSuits {
			cards = append(cards, value+" of "+suits)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// d deck -> receiver (where d is reference{like this in js} and deck is type)
// Meaning -> Any variable of type "deck" now gets accesss to the "print" method

// Format:- func (receiver(s))func_name(argument(s)) (return type(s))

//Multiple value return:
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// type conversion:- []byte (string)

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//convert byte slice(bs) to deck type -> Reverse process
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		//Generate random number between 0 and length of deck:-
		newPosition := r.Intn(len(d) - 1)
		// swap elements:-
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
