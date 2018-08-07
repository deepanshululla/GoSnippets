package main
import "fmt"

/*
The issue with this is it doesn't return
any results because main thread finishes before the
other go routine
Use wait groups to ask main thread to wait for the other go routine execution

 */
func main() {
	go foo1()
	go bar1()
}

func foo1() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar1() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
}

