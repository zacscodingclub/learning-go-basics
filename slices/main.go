package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	test()
}

func first() {
	mySlice := []string{"a", "b", "c", "g", "l", "m"}

	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])
	fmt.Println(mySlice[2])
	fmt.Println("myString"[2])
}

func second() {
	mySlice := make([]int, 0, 5)

	fmt.Println("----------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("----------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value:", mySlice[i])
	}
}

func appending() {
	mySlice := []int{1, 2, 3, 4, 5}
	otherSlice := []int{6, 7, 8, 9}

	mySlice = append(mySlice, otherSlice...)

	fmt.Println(mySlice)
}

func deleting() {
	mySlice := []string{"Monday", "Tuesday"}
	otherSlice := []string{"Wednesday", "Thursday", "Friday"}

	mySlice = append(mySlice, otherSlice...)
	fmt.Println(mySlice)

	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println(mySlice)
}

func test() {
	var nonce [24]byte
	fmt.Println(nonce)
	io.ReadFull(rand.Reader, nonce[:])
	fmt.Println(nonce)
}
