package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {

	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Club", "Hearts"}
	cardsValue := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Gulam", "Queen", "King"}
	for _, suits := range cardSuits {
		for _, value := range cardsValue {
			cards = append(cards, value+" of "+suits)

		}
	}
	return cards

}
func newNumList() deck {

	numList := deck{}
	firstNum := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	lastNum := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	for _, fnum := range firstNum {
		for _, lnum := range lastNum {
			numList = append(numList, fnum+lnum)
		}

	}
	return numList
}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToHDD(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 06666)
}

func newDeckFromFile(filename string) deck {

	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	// fmt.Println(string(bs))
	s := strings.Split(string(bs), ",")
	return deck(s)
	// return ("Hello")
}

func (d deck) suffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
