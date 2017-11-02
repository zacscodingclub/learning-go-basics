package pipeline

import (
	"fmt"
	"math/big"
)

func p(c *big.Int) {
	fmt.Println(c)
}
func Main() {
	in := genFact()

	f := factorial(in)
	for n := range f {
		p(n)
	}
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
func genFact() <-chan int64 {
	out := make(chan int64)
	go func() {
		for i := 0; i < 15; i++ {
			for j := 3; j < 20; j++ {
				out <- int64(j)
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int64) <-chan *big.Int {
	out := make(chan *big.Int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int64) *big.Int {
	total := big.NewInt(1)
	for i := n; i > 0; i-- {
		total = mul(total, big.NewInt(i))
	}

	return total
}

func mul(x, y *big.Int) *big.Int {
	return big.NewInt(0).Mul(x, y)
}
