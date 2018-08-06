package main

import ("fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)

//create a new deck type
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "King", "Queen", "Jack", "Joker"}
	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}
	return cards;
}

func (d deck) printEachCard() {
	fmt.Println("printing cards from GO method")
	for i, card := range d {
		fmt.Println(i, card)
	}
}
func (d deck) toString() string{
	return strings.Join([]string(d),",")
}

func deal(d deck,handsize int) (deck,deck){
	return d[:handsize],d[handsize:]
}

func (d deck) saveToFile(fileName string) error{
	return ioutil.WriteFile(fileName,[]byte(d.toString()),0666);
}

func newDeckFromFile(filename string) deck{
	cardBytes,err:=ioutil.ReadFile(filename)
	if (err!=nil){
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - Log the error and entirely quit the program

		fmt.Println("error has occured while reading")
		fmt.Println(err)
		os.Exit(1)
	}
	s:=strings.Split(string(cardBytes),",")
	return deck(s)

}

func (d deck) shuffle(){
	source:=rand.NewSource(time.Now().UnixNano())
	r:=rand.New(source)
	for i := range d{
		newPosition := r.Intn(len(d)-1)
		d[i],d[newPosition]=d[newPosition],d[i]

	}
}