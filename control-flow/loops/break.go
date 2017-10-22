package loops

import "fmt"

func Break() {
	i := 0
	for {
		fmt.Println(i)
		if i >= 10 {
			break
		}

		i++
	}
}
