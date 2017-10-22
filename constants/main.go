package main

import "fmt"

const p string = "Death & Taxes"

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
)

func main() {
	const q = 42
	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println(KB)
	fmt.Println(MB)
}
