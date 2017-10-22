package main

import (
	"fmt"
	"runtime"
)

func trace() {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])

	fmt.Printf(f.Name())
}

func printAddresses(val int) {
	fmt.Printf("%p\n", &val)
	fmt.Println(&val)
}

func zero(x int) {
	trace()
	printAddresses(x)
	x = 0
}

func realZero(x *int) {
	*x = 0
}

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b = &a // pointer integer type

	fmt.Println(b)
	fmt.Println(*b) // deregisters value in memory address

	*b = 42
	// a is now equal to 42 since we changed the value in that memory bucket
	fmt.Println(a)

	x := 5
	trace()
	printAddresses(x)
	zero(x)
	fmt.Println(x)
	realZero(&x)
	fmt.Println(x)
}
