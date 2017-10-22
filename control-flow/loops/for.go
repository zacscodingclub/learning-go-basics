package loops

import "fmt"

func For(n int) {
	for i := 0; i <= n; i++ {
		fmt.Println(i)
	}
}