package main

import (
	"fmt"
	"sync"
)

var workerID, publisherID int

func main() {
	in := gen()
	f1 := factorial(in)
	f2 := factorial(in)
	f3 := factorial(in)
	f4 := factorial(in)
	f5 := factorial(in)
	f6 := factorial(in)
	f7 := factorial(in)
	f8 := factorial(in)
	f9 := factorial(in)
	y := 0
	for n := range merge(f1, f2, f3, f4, f5, f6, f7, f8, f9) {
		fmt.Println(y, "\t", n)
		y++
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func merge(chans ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(chans))

	for _, c := range chans {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func publisher(out chan string) {
	publisherID++
	thisID := publisherID
	dataID := 0
	for {
		dataID++
		fmt.Printf("publisher %d is pushing data\n", thisID)
		data := fmt.Sprintf("Data from publisher %d.  Data %d", thisID, dataID)
		out <- data
	}
}

func workerProcess(in <-chan string) {
	workerID++
	thisID := workerID
	for {
		fmt.Printf("%d: waiting for input...\n", thisID)
		input := <-in
		fmt.Printf("%d: input is %s\n", thisID, input)
	}
}

/*
Is this fan out?
Sort of.  We are fanning out by launching multiple go routines
that are publishing a message into our channel.  Golang blog says,
"Fan out means you have multiple functions reading from the same
channel."

Is this fan in?
No. To "fan in", it must read from multiple channels, but this only
reads from one.  Golang blog describes fan in: "A function can read
from multiple inputs and proceed until all are closed by
multiplexing the input channels onto a single channel that's closed
when all the inputs are closed."
*/
