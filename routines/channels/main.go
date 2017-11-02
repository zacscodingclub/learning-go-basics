package channels

import (
	"fmt"
	"sync"
	"time"
)

func Example() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i // put into the channel
		}
	}()

	go func() {
		for {
			fmt.Println(<-c) //  take from the channel
		}
	}()

	time.Sleep(time.Second)
}

func RangeClause() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i // put into the channel
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

func NToOne() {
	c := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
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

func SemaphoreOne() {
	c := make(chan int)

	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(c)
		close(done)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

func SemaphoreTwo() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			done <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()
	for n := range c {
		fmt.Println(n)
	}
}

func OneToN() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10000; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i++ {
		go func() {
			for n := range c {
				fmt.Println("channel:", i, n)
			}

			done <- true
		}()
	}

	for i := 0; i < n; i++ {
		<-done
	}
}

func ArgsAndReturns() {
	c := incrementor()
	cSum := puller(c)

	for n := range cSum {
		fmt.Println(n)
	}
}

func incrementor() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
