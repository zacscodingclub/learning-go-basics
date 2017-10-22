package exercises

import "fmt"

func HelloWorld() {
	fmt.Println("Hello World")
}

func HelloName(n string) {
	fmt.Println("Hello", n)
}

func Hello() {
	var name string
	fmt.Println("What's your name?")
	fmt.Scan(&name)

	HelloName(name)
}
