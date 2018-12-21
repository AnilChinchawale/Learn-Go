package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades,But got %v", d[0])
	}
	if d[len(d)-1] != "King of Hearts" {
		t.Errorf("Expected last card of King of Hearts, But we got %v", d[len(d)-1])
	}
	if d[len(d)-2] != "Queen of Hearts" {
		t.Errorf("Expected 2nd last card of Queen of Hearts, But we got %v", d[len(d)-2])
	}
}

func TestSaveToHDDAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	newfilename := "_olddecktesting"
	os.Remove(filename)
	deck := newDeck()
	deck.saveToHDD(filename)
	loaded := newDeckFromFile(filename)
	if len(loaded) != 52 {
		t.Errorf("Expected 52 cards in deck,But got %v", len(loaded))
	}

	os.Rename(filename, newfilename)
	os.Remove(filename)
}
