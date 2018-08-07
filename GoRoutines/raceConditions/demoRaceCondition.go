package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
)

var wg sync.WaitGroup
var counter int

func main(){
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string){
	// choose a random seed so the random numbers are generated differently in each run
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}
/*
sjc-mpfgm:raceConditions dlulla$ go run -race demoRaceCondition.go
==================
WARNING: DATA RACE
Write at 0x0000011da620 by goroutine 6:
  main.incrementor()
      /Users/dlulla/mycode/GoProjects/src/github.com/deepanshululla/GoSnippets/GoRoutines/raceConditions/demoRaceCondition.go:28 +0x152

Previous write at 0x0000011da620 by goroutine 7:
  main.incrementor()
      /Users/dlulla/mycode/GoProjects/src/github.com/deepanshululla/GoSnippets/GoRoutines/raceConditions/demoRaceCondition.go:28 +0x152

Goroutine 6 (running) created at:
  main.main()
      /Users/dlulla/mycode/GoProjects/src/github.com/deepanshululla/GoSnippets/GoRoutines/raceConditions/demoRaceCondition.go:15 +0x74

Goroutine 7 (running) created at:
  main.main()
      /Users/dlulla/mycode/GoProjects/src/github.com/deepanshululla/GoSnippets/GoRoutines/raceConditions/demoRaceCondition.go:16 +0xa1
==================

 */