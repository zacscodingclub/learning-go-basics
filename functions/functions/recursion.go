package functions

func Factorial(n int) int {
	if n < 2 {
		return n
	}
	return n * Factorial(n-1)
}
