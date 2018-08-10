package main

import "fmt"

/*
The optional <- operator specifies the channel direction, send or receive.
If no direction is given, the channel is bidirectional.
https://golang.org/ref/spec#Channel_types

chan T          // can be used to send and receive values of type T
chan<- float64  // can only be used to send float64s
<-chan int      // can only be used to receive ints

*/


func main()  {
	c := incrementor2()
	cSum := puller2(c)
	for n := range cSum {
		fmt.Println(n)
	}
}
// can only be used to return ints
func incrementor2() <-chan int  {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}
// receive ints and the output them
func puller2(c <-chan int ) <-chan int  {
	out2 := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out2 <- sum
		close(out2)
	}()
	return out2
}
