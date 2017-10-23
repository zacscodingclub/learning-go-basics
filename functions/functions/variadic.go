package functions

func Average(nums ...float64) float64 {
	var total float64
	for _, v := range nums {
		total += v
	}

	return total / float64(len(nums))
}
