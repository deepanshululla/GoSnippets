package main

import (
	"sync"
	"math/rand"
	"fmt"
)

const numFactorials = 100
const rdLimit = 20


func main() {
	var wgMain sync.WaitGroup;
	wgMain.Add(numFactorials)
	factorialWait(&wgMain)
	wgMain.Wait()
}


func factorialWait(wgMain *sync.WaitGroup){
	var w sync.WaitGroup
	rand.Seed(42)

	w.Add(numFactorials + 1)
	for j := 1; j <= numFactorials; j++ {
		go func() {
			f := rand.Intn(rdLimit)
			w.Wait()
			total := factWait(f)
			fmt.Println(f, total)
			(*wgMain).Done()
		}()
		w.Done()
	}
	fmt.Println("All done with initialization")
	w.Done()
}

func factWait(n int) int{
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return int(total)
}

/*
CHALLENGE #1:
-- Change the code above to execute 100 factorial computations concurrently and in parallel.
-- Use the "pipeline" pattern to accomplish this using wait groups

POST TO DISCUSSION:
-- What realizations did you have about working with concurrent code when building your solution?
-- eg, what insights did you have which helped you understand working with concurrency?
-- Post your answer, and read other answers, here: https://goo.gl/uJa99G
*/


