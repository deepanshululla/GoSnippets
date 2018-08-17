package main

import "fmt"

func main() {
	for n:= range (factorialChan(gen(10))){
		fmt.Println(n)
	}
}

func gen(num int) chan int{
	out:=make(chan int)
	go func() {
		for i:=0;i<num+1;i++{
			out<-i
		}
		close(out)
	}()
	return out
}



func factorialChan(c chan int) chan int64 {
	out := make(chan int64)
	go func() {

		for i:=range c{
			total :=fact(i)
			out <- total
		}

		close(out)
	}()
	return out
}

func fact(n int) int64{
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return int64(total)
}

/*
CHALLENGE #1:
-- Change the code above to execute 100 factorial computations concurrently and in parallel.
-- Use the "pipeline" pattern to accomplish this

POST TO DISCUSSION:
-- What realizations did you have about working with concurrent code when building your solution?
-- eg, what insights did you have which helped you understand working with concurrency?
-- Post your answer, and read other answers, here: https://goo.gl/uJa99G
*/

