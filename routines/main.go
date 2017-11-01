package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

// init like normal initialization
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // User all cours
}

func main() {
	wg.Add(2)
	go incrementer("foo")
	go incrementer("bar")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementer(s string) {
	for i := 0; i < 100; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter", counter)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 100; i++ {
		fmt.Println("Bar:", i)
	}
	wg.Done()
}
