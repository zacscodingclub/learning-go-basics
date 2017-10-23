package functions

import "fmt"

func Expression() {
	// anonymous function && closure
	greeting := func() {
		fmt.Println("Hello World")
	}

	greeting()
}

func MakeGreeter() func() string {
	// closure
	return func() string {
		return "Hallo World!"
	}
}
