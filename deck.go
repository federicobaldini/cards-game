package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type card struct {
	value string
	suit  string
}

// Create a new type of "deck"
// which is a slice of strings
type deck []card

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Height", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			new_card := card{value, suit}
			cards = append(cards, new_card)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card.value+" of "+card.suit)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	string_deck := []string{}
	for _, card := range d {
		string_deck = append(string_deck, card.value+" of "+card.suit)
	}
	return strings.Join([]string(string_deck), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	ss := strings.Split(string(bs), ",")
	new_deck := []card{}
	for _, string_item := range ss {
		string_card := strings.Split(string_item, " of ")
		new_deck = append(new_deck, card{string_card[0], string_card[1]})
	}
	return deck(new_deck)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixMicro())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
