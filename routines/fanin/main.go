package fanin

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Main() {
	c := fanIn(boring("joe"), boring("ann"))

	for n := range c {
		fmt.Println(n)
	}
	fmt.Println("You're both boring; I'm leaving!")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 11; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

/*
code source:
Rob Pike
https://talks.golang.org/2012/concurrency.slide#25

FAN OUT
Multiple functions reading from the same channel until that channel is closed
FAN IN
A function can read from multiple inputs and proceed until all are closed by
multiplexing the input channels onto a single channel that's closed when
all the inputs are closed.
PATTERN
there's a pattern to our pipeline functions:
-- stages close their outbound channels when all the send operations are done.
-- stages keep receiving values from inbound channels until those channels are closed.
source:
https://blog.golang.org/pipelines
*/

func MainTwo() {
	in := gen(2, 3)

	// FAN OUT
	// distribute the sq work across two go routines that both read from in.
	c1 := fanOut(in)
	c2 := fanOut(in)

	// FAN IN
	// consume the merged output ffrom multiple channels
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func fanOut(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(chans ...chan int) chan int {
	fmt.Printf("TYPE OF CHANS: %T\n", chans)
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(chans))

	for _, c := range chans {
		go func(c chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
