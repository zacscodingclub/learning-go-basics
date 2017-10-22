package exercises

import "fmt"

func Remainder() {
	var small, large int
	fmt.Println("Enter a small number:")
	fmt.Scan(&small)
	fmt.Println("Enter a large number:")
	fmt.Scan(&large)
	fmt.Println("The remainder is:", large%small)
}
