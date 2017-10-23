package functions

func Visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func Filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}

	return xs
}
