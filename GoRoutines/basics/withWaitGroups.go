package main

import (
	"sync"
	"fmt"
)

var wg1 sync.WaitGroup;

func main(){
	wg1.Add(2)
	go foo3()
	go bar3()
	wg1.Wait()

}
func foo3() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
	wg1.Done()
}

func bar3() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
	wg1.Done()
}
