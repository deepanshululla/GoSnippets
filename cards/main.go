package main

import "fmt"

func main(){
	//cards:=deck{"Ace of spades",newCard()}
	//cards:=newDeck()
	//cards=append(cards,newCard())
	//fmt.Println(cards)
	//for i,card:=range cards{
	//	fmt.Println(i,card);
	//}
	//cards.printEachCard()
	//hand,remainingDeck:=deal(cards,5)
	//fmt.Println(hand)
	//fmt.Println(remainingDeck)
	//fmt.Println("Convert to string")
	//fmt.Println(remainingDeck.toString())
	//cards.saveToFile("my_cards")
	cards:=newDeckFromFile("my_cards")
	fmt.Println(cards)
	cards.shuffle()
	fmt.Println(cards)
}




func newCard() string{
	return "5 of Hearts"
}