package main

import "fmt"

func main()  {
	c:=make(chan int);
	// basically keep adding to channel
	// and then close it
	go func() {
		for i:=0;i<10;i++{
			c<-i
		}
		close(c)
	}()

	// placing close here causes a race condition because
	// it will close the channel before the next value to channel is sent
	//close(c)


	// now iterate over the values added
	// this is going to wait from the values received from the channel
	for n:=range c{
		fmt.Println(n)
	}
}

