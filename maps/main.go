package main

import (
	"fmt"
)

func basics(){

	//option1
	//var colors map[string]string

	//option2
	//colors:=map[string]string{
	//	"red":"ddfsfs",
	//	"blue":"bdjfbsk",
	//}

	//option 3
	colors:=make(map[string]string)
	colors["white"]="#ffffff"
	colors["red"]="#ff0000"

	fmt.Println(colors)
	delete(colors,"red")
	fmt.Printf("%+v\n",colors)
}

func iterateOverMap(){
	colors:=map[string]string{
		"green":"#00ff00",
		"blue":"#0000ff",
		"red":"#ff0000",
	}
	//for key,value := range colors {
	//	fmt.Println(key,value)
	//}
	iterate(colors)

}

func iterate(colors map[string]string){
	for _,value := range colors {
		fmt.Println(value)
	}
}

func main() {

	iterateOverMap()
}
