package main

import "fmt"

func main() {
	arr:=make([]int,10)
	for index:= range arr{
		arr[index]=index+1;
	}
	fmt.Println(arr);

	for index:=range arr{
		if (arr[index]%2==0){
			fmt.Println("even")
		}else{
			fmt.Println("odd")
		}
	}
}
