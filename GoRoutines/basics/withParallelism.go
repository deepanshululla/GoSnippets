package main

import (
	"sync"
	"runtime"
	"fmt"
	"time"
)
var wg3 sync.WaitGroup;

/*
Before golang 1.5 we need to specify runtime.GOMAPXPROCS to say to the compiler
to not limit yourself to one CPU.
Use multiple cores .otherise go would just use one cpu core.
However after 1.5 this has become redundant
 */
func init(){
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main()  {
	wg3.Add(2)
	go foo5()
	go bar5()
	wg3.Wait()
}

func foo5() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(300 * time.Millisecond)
	}
	wg2.Done()
}

func bar5() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(300 * time.Millisecond)
	}
	wg2.Done()
}
