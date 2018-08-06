package main

import "fmt"


func main() {
	var i,j,k=5,3,"foo"
	fmt.Println("hello world Deepanshu")
	fmt.Print(i,j," ",k,"\n")
	fmt.Printf("i is of type %T\n",i)
	fmt.Printf("j is of type %T\n",j)
	fmt.Printf("k is of type %T\n",k)
	if (i>2){
		fmt.Printf("i which is %d is greater than 2 \n",i);
	}
	greeting:="hi there"
	fmt.Println([]byte(greeting))

}
/*
This is a multiline comment

 */