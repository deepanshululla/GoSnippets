package main

/*
Testing Go routine inside a Go routine
 */
import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo2()
	go bar2()
	wg.Wait()
}

func foo2() {
	wg.Add(1)
	go someFunc()
	for i := 0; i < 45; i++ {
		fmt.Println("Foo2:", i)
	}

	wg.Done()
}
func someFunc(){
	fmt.Println("Voila")
	wg.Done()
}

func bar2() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar2:", i)
	}
	wg.Done()
}

