package main

import "fmt"

func main(){
	//var card string="Ace of Spades"
	card:="Ace of spades"
	//both of above statements do the same thing..declare new variable and define it will be a string
	//:= is asking compiler to figure out the type. It is always used for now variables and not reasigning
	//card="five of diamonds"
	//card=5---> this will generate an error since we have defined card to be of type string

	card=newCard()
	fmt.Println(card)
	arraysInGo()
}

func arraysInGo(){
	/*
	Go has two types of arrays fixed length
	and dynamic arrays(called slices)
	All the elements in slice shoudl be of same type
	 */
	cards:=[]string{"Ace of spades","Five of diamonds","wtf"};
	cards=append(cards,"9 of hearts")
	fmt.Println(cards);
	for i,card:=range cards{
		fmt.Println(i,card)
	}
}


func newCard() string{
	return "5 of Spades"
}