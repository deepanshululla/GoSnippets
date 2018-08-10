package main

import (
	"sync"
	"fmt"
)

func main(){
	c := make(chan int)

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}


}
/*
sjc-mpfgm:NTo1 dlulla$ go run -race raceConditionWithChannel.go
0
==================
0
WARNING: DATA RACE
Write at 0x00c42009601c by goroutine 8:
  internal/race.Write()
      /usr/local/Cellar/go/1.10.3/libexec/src/internal/race/race.go:41 +0x38
1
  sync.(*WaitGroup).Wait()
      /usr/local/Cellar/go/1.10.3/libexec/src/sync/waitgroup.go:127 +0xf3
  main.main.func3()
      /Users/dlulla/mycode/GoProjects/src/github.com/deepanshululla/GoSnippets/GoRoutines/channels/NTo1/raceConditionWithChannel.go:30 +0x38

Previous read at 0x00c42009601c by goroutine 6:
  internal/race.Read()
      /usr/local/Cellar/go/1.10.3/libexec/src/internal/race/race.go:37 +0x38
2
  sync.(*WaitGroup).Add()
      /usr/local/Cellar/go/1.10.3/libexec/src/sync/waitgroup.go:70 +0x16e
  main.main.func1()
      /Users/dlulla/mycode/GoProjects/src/github.com/deepanshululla/GoSnippets/GoRoutines/channels/NTo1/raceConditionWithChannel.go:14 +0x45

Goroutine 8 (running) created at:
3
  main.main()
      /Users/dlulla/mycode/GoProjects/src/github.com/deepanshululla/GoSnippets/GoRoutines4
5
/channels/NTo1/raceConditionWithChannel.go:29 +0x107

Goroutine 6 (running) created at:
6
  main.main()
      /Users/dlulla/mycode/GoProjects/src/github.com/deepanshululla/GoSnippets/GoRoutines/channels/NTo1/raceConditionWithChannel.go:13 +0xaf
==================

 */

