package functions

import (
	"fmt"
	"math"
	"strconv"
)

func Half(n int) (float64, bool) {
	div := float64(n) / 2
	divByTwo := n%2 == 0
	return div, divByTwo
}

func HalfExpression(n int) (float64, bool) {
	half := func(x int) (float64, bool) {
		div := float64(x) / 2
		divByTwo := x%2 == 0
		return div, divByTwo
	}

	return half(n)
}

func FindMax(args ...float64) float64 {
	var largest float64
	for _, v := range args {
		if v > largest {
			largest = v
		}
	}

	return largest
}

func BoolReturn() bool {
	return (true && false) || (false && true) || !(false && false)
}

// https://projecteuler.net/problem=56

func sumDigits(val string) float64 {
	vals := []rune(val)
	var sum float64

	for _, v := range vals {
		intValue := v - '0'
		if intValue == -2 {
			break
		}
		sum += float64(intValue)
	}

	return sum
}

func MaxDigitalSum() (float64, float64) {
	var largestSum, first, second, i, j float64

	for i = 99; i > 0; i-- {
		for j = 99; j > 0; j-- {
			exp := strconv.FormatFloat(math.Pow(i, j), 'f', 6, 64)

			sum := sumDigits(exp)
			if sum > largestSum {
				largestSum, first, second = sum, i, j
			}
		}
	}

	fmt.Println("Largest:", largestSum)
	fmt.Println("Values:", first, "&", second)
	return first, second
}
