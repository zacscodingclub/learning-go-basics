package main

import (
	"fmt"

	"github.com/zacscodingclub/learning-go-basics/scope/vis"
)

var x = 42

func main() {
	fmt.Println(x)
	foo()
	fmt.Println(vis.MyName)
	vis.PrintVar()
}

func foo() {
	fmt.Println(x)
}
