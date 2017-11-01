package main

import (
	"fmt"
	"math"
	"sort"
)

type people []person

type person struct {
	name string
}

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	area() float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

// value receiver
func (z Circle) area() float64 {
	return math.Pi * z.radius * z.radius
}

// pointer reciever
func (z *Circle) pArea() float64 {
	return math.Pi * z.radius * z.radius
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i].name[0] < p[j].name[0] }

func main() {
	// s := Square{10}
	// info(s)

	// c := Circle{3}
	// info(c)

	// Sort slice of string
	// try different methods
	// then try reverse
	// studyGroup := people{{"Zeno"}, {"John"}, {"Al"}, {"Jenny"}}
	//sort.Sort(studyGroup) // normal
	// sort.Sort(sort.Reverse(studyGroup))
	p := []string{"Zeno", "Zac", "Roxanne", "John", "Al", "Jenny"}
	// sort.Strings(p) // normal
	// sort.StringSlice(p).Sort()
	sort.Sort(sort.StringSlice(p))
	// sort.Sort(sort.Reverse(sort.StringSlice(p))) // reverse

	// fmt.Println(p)
	// sort.Strings(p)
	// fmt.Println(p)
	// n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	// sort.Int(s) // normal
	// sort.Sort(sort.Reverse(sort.IntSlice(n))) // reverse
	fmt.Println(p)
}

/*
Polymorphism is the ability to write code that can take on different
behaviour through the implementation of types.  Once a type implements
an interface, an entire world of functionality can be opened up to
values of that type.
-
*/
