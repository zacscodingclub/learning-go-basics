package loops

import "fmt"

func Continue() {
	i := 0
	for {
		i++
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
}
