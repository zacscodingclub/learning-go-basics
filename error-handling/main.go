package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

// convention to create all the errors at the top of the file
// then they can be used throughout.
// nothing "exceptional" about an error, it's just a string
var ErrNegativeMath = errors.New("norgate math: square root of negative number")

func main() {
	fmt.Println(sqrt(-2))
	fmt.Println(sqrt(99))
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeMath
	}

	return math.Sqrt(f), nil
}

func example1() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("Error:", err)

		// log.Println("Error:", err)
		// # command-line-arguments
		// ./main.go:15:3: undefined: log

		// log.Fatalln(err)
		// 2017/11/06 20:55:03 open no-file.txt: no such file or directory
		// exit status 1

		panic(err)
		/*
			panic: open no-file.txt: no such file or directory

			goroutine 1 [running]:
			main.example1()
			/Users/zb/Development/golang/src/github.com/zacscodingclub/learning-go-basics/error-handling/main.go:20 +0x63
			main.main()
			/Users/zb/Development/golang/src/github.com/zacscodingclub/learning-go-basics/error-handling/main.go:8 +0x20
			exit status 2
		*/
	}
}
