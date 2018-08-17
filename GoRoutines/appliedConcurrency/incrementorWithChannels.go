package main

import "fmt"

func incrementorChan(s string) chan int{
	out:=make(chan int)
	go func() {
		for i:=0;i<20;i++{
			out<-1
			fmt.Println(s,i)
		}
		close(out)
	}()
	return out
}


func pullerChan(c chan int) chan int{
	out:=make(chan int)
	go func() {
		var sum int
		for n:= range c{
			sum+=n
		}
		out<-sum
		close(out)
	}()
	return out
}

func main()  {
	c1 := incrementorChan("Foo:")
	c2 := incrementorChan("Bar:")
	c3 := pullerChan(c1)
	c4 := pullerChan(c2)
	fmt.Println("Final Counter:", <-c3+<-c4)
}
