package functions

import "fmt"

func print(i interface{}) {
	fmt.Println("during:", i)
}

func ChangeMe(z *int) {
	print(z)
	print(*z)
	*z = 24
	print(z)
	print(*z)
}
