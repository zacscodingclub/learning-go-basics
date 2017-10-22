package main

import "fmt"

const metersToYards float64 = 1.09361

func showingAddress() {
	a := 12

	fmt.Println("a: ", a)
	fmt.Println("a's memory address", &a)
	fmt.Printf("%d \n", &a)
}

func convertYardsToMeters() {
	var meters float64
	fmt.Print("Enter meters ran: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, "yards.")
}

func main() {
	showingAddress()
	convertYardsToMeters()
}
