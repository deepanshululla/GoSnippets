package main

import (
	"sync"
	"fmt"
	"time"
)

var wg2 sync.WaitGroup;

func main(){
	wg2.Add(2)
	go foo4()
	go bar4()
	wg2.Wait()
}

func foo4() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(300 * time.Millisecond)
	}
	wg2.Done()
}

func bar4() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(300 * time.Millisecond)
	}
	wg2.Done()
}
