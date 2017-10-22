package exercises

import (
	"fmt"
)

func FizzBuzz() {
	for i := 0; i <= 100; i++ {
		fizz := "Fizz"
		buzz := "Buzz"

		if i%15 == 0 {
			fmt.Println(fizz + buzz)
		} else if i%3 == 0 {
			fmt.Println(fizz)
		} else if i%5 == 0 {
			fmt.Println(buzz)
		} else {
			fmt.Println(i)
		}
	}
}
