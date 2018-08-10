package main

import (
	"sync"
	"time"
	"math/rand"
	"sync/atomic"
	"fmt"
)

var wg sync.WaitGroup
var counter int64

func main(){
	//wg.Add(2)
	//go incrementor("Foo:")
	//go incrementor("Bar:")
	//wg.Wait()
	//fmt.Println("Final Counter:", counter)

	wg.Add(2)
	go incrementor2("Foo:")
	go incrementor2("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string){
	for i:=0;i<20;i++{
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", atomic.LoadInt64(&counter)) // access without race
	}
	wg.Done()
}

func incrementor2(s string){
	// if we do counter<20, this is not an atomic operation and causes a race condition
	//for ;counter<20;{
	for ;atomic.LoadInt64(&counter)<20;{
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, "Counter:", atomic.LoadInt64(&counter)) // access without race
	}
	wg.Done()
}

/*
In this neither incrementor1 nor incrementor2 have race conditions
This is because all the read and write operations are atomic operations.
In incrementor2 even if one of the operations is non atomic that could lead to race conditions

 */
