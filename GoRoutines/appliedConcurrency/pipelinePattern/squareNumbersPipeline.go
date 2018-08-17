package main

import "fmt"

func gen(nums ...int) <-chan int{
	out:=make(chan int)
	go func() {
		for _,n:=range nums{
			out<-n
		}
		close(out)
	}()
	return out
}
// note the ,_ sign indicates receive only channel and is optional
func sq(in <-chan int) <-chan int{
	out:=make(chan int)
	go func() {
		for n:=range in{
			out<-n*n
		}
		close(out)
	}()
	return out
}

func main() {
	//set up the pipeline
	c := gen(2, 3)
	out:= sq(c)
	//Consume the output.
	for n:=range out{
		fmt.Println(n)
	}
	// refactor the code, setup the pipeline and consume the output
	for n:=range (sq(sq(gen(2, 3)))){
		fmt.Println(n)
	}

}
