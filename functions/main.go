package main

import (
	f "github.com/zacscodingclub/learning-go-basics/functions/functions"
)

func main() {
	// f.Greet("Zac", "Bob")
	// fmt.Println(f.GreetReturn("Bill", "Joe"))
	// fmt.Println(f.MultipleReturn("Steve", "Julie"))

	// defer f.Hello()
	// n := f.Average(43, 55, 39, 88, 12, 66, 99)
	// fmt.Println(n)

	// data := []float64{32, 43, 56, 99, 13, 55, 75}
	// // spreads this into a float 64 type, currently it's a slice
	// a := f.Average(data...)
	// fmt.Println(a)

	// f.Expression()
	// greet := f.MakeGreeter()
	// fmt.Println(greet())

	// double := func(n int) {
	// 	fmt.Println(n * 2)
	// }
	// f.Visit([]int{1, 2, 3, 4}, func(n int) {
	// 	double(n)
	// })

	// vals := f.Filter([]int{1, 2, 3, 4}, func(n int) bool {
	// 	return n%2 == 0
	// })
	// fmt.Println(vals)
	// fmt.Println(f.Factorial(5))

	// age := 33
	// fmt.Println("before:", &age)
	// f.ChangeMe(&age)
	// fmt.Println("after:", &age)
	// fmt.Println("after:", age)

	// func() {
	// 	fmt.Println("IIFE")
	// }()

	// fmt.Println(f.Half(1))
	// fmt.Println(f.Half(2))

	// fmt.Println(f.HalfExpression(3))
	// fmt.Println(f.HalfExpression(16))

	// fmt.Println(f.FindMax(43, 55, 39, 88, 12, 66, 99))
	// fmt.Println(f.FindMax(32, 43, 56, 99, 13, 55, 75))

	// fmt.Println(f.BoolReturn())

	f.MaxDigitalSum()
}

// func (receivers) name(parameters type) return type {}
