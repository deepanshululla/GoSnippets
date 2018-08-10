package main

import (
	"time"
	"math/rand"
	"fmt"
	"sync"
)
var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main(){
	// one without race condition
	//wg.Add(2)
	//go incrementor("Foo:")
	//go incrementor("Bar:")
	//wg.Wait()
	//fmt.Println("Final Counter:", counter)

	// one with race condition
	//wg.Add(2)
	//go incrementor2("Foo:")
	//go incrementor2("Bar:")
	//wg.Wait()
	//fmt.Println("Final Counter:", counter)

	// this just creates an infinite loop
	// though it doesn't create a race condition
	wg.Add(2)
	go incrementor3("Foo:")
	go incrementor3("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

/*
Note that just because you are using mutex doesn't mean you avoid race conditions at all times
There is no race condition in incrementor but there is one in incrementor2
 */
func incrementor(s string){
	for i:=0;i<20;i++{
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Println(s, i, "Counter:", counter)
		mutex.Unlock()
	}
	wg.Done()
}

func incrementor2(s string){
	for ;counter<20;{
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Println(s, "Counter:", counter)
		mutex.Unlock()
	}
	wg.Done()
}

func incrementor3(s string){
	for {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		mutex.Lock()
		if (counter>20) {
			break
		}
		counter++
		fmt.Println(s, "Counter:", counter)
		mutex.Unlock()
	}
	wg.Done()
}
/*
In the incrementor2() method , so the race condition happens at lines 45 and 48 makes sense because
these are the two places where counter is trying to be accessed.


sjc-mpfgm:mutex dlulla$ go run -race mutexDemo.go
==================
WARNING: DATA RACE
Write at 0x0000011da620 by goroutine 6:
  main.incrementor2()
      /Users/dlulla/mycode/GoProjects/src/github.com/deepanshululla/GoSnippets/GoRoutines/mutex/mutexDemo.go:48 +0xa2

Previous read at 0x0000011da620 by goroutine 7:
  main.incrementor2()
      /Users/dlulla/mycode/GoProjects/src/github.com/deepanshululla/GoSnippets/GoRoutines/mutex/mutexDemo.go:45 +0x1e8

Goroutine 6 (running) created at:
  main.main()
      /Users/dlulla/mycode/GoProjects/src/github.com/deepanshululla/GoSnippets/GoRoutines/mutex/mutexDemo.go:23 +0x74

Goroutine 7 (running) created at:
  main.main()
      /Users/dlulla/mycode/GoProjects/src/github.com/deepanshululla/GoSnippets/GoRoutines/mutex/mutexDemo.go:24 +0xa1
==================

 */