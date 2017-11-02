package main

/*
Channels Resources
https://golang.org/doc/effective_go.html#concurrency



*/
import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/zacscodingclub/learning-go-basics/routines/fanin"
)

var wg sync.WaitGroup
var counter int64

// init like normal initialization
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // User all cours
}

func main() {
	fanin.MainTwo()
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
