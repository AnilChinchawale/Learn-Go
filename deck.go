package main

import (
	"fmt"
	"io/ioutil"
	"strings"
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
	cardsValue := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Gulam", "Queen", "King"}
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
