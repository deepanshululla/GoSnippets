package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T){
	d:=newDeck()
	if(len(d)!=52){
		t.Errorf("Expected deck length of size 52, got length %v\n",len(d))
	}
	if d[0]!="Ace of Spades" || d[len(d)-1]!="Joker of Clubs"{
		t.Errorf("Error in ordering of cards. First card should be Ace " +
			"of Spades but is %v and last card should be Joker of Clubs but is %v ",d[0],d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting");
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck:=newDeckFromFile("_decktesting")
	if (len(loadedDeck)!=len(deck) || len(loadedDeck)!=52){
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}
	os.Remove("_decktesting");
}