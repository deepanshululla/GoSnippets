package main

import (
	"fmt"
)
/*
Basically in this we have created two channels
one to store the data values and second
to store the values as in the signal
We  are storing n*10 values in the c channel
and we are storing the notify in done channel
 */
func main() {

	n := 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			done <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

