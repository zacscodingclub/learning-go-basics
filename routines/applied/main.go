package applied

import (
	"fmt"
)

func Deadlocked() {
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}

/*
This results in a deadlock.
Can you determine why?
And what would you do to fix it?
*/

func NotDeadlocked() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}

func OnlyPrintsZero() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	fmt.Println(<-c)
}

// Why does it only print zero?
// We are only popping off the first item in the channel

func PrintAllChannel() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

func BasicFactorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}

	return total
}

// Use go routines and channels to calculate factorial
func ChannelFactorialOne(n int) int {
	total := 1
	c := make(chan int)

	go func() {
		for i := n; i > 0; i-- {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		total *= n
	}

	return total
}

func ChannelFactorialTwo(n int) int {
	c := make(chan int)

	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		c <- total
		close(c)
	}()

	return <-c
}
