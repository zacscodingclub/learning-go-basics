package main

import "fmt"

func wrapper() func() int {
	x := 0
	fmt.Println("From insider the wrapper")

	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
