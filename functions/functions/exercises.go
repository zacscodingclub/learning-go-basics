package functions

import (
	"fmt"
	"math/big"
	"time"
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

func SumDigits(val string) int64 {
	vals := []rune(val)
	var sum int64

	for _, v := range vals {
		intValue := v - '0'

		sum += int64(intValue)
	}

	return sum
}

func MaxDigitalSum() (int64, int64) {
	start := time.Now()
	var largestSum, first, second, i, j int64

	for i = 99; i > 0; i-- {
		for j = 99; j > 0; j-- {

			exp := new(big.Int).Exp(big.NewInt(i), big.NewInt(j), nil)

			sum := SumDigits(exp.String())
			if sum > largestSum {
				largestSum, first, second = sum, i, j
			}
		}
	}

	t := time.Now()
	fmt.Println("Largest:", largestSum)
	fmt.Println("Values:", first, "&", second)
	fmt.Println("Time:", t.Sub(start))
	return first, second
}
