package loops

import "fmt"

func While() {
	i := 0
	for i < 10 {
		if i > 0 {
			fmt.Println(i+1, " strange loops")
		} else {
			fmt.Println(i+1, " strange loop")
		}
		i++
	}
}
