package loops

import "fmt"

func wassup(n string) {
	fmt.Println("Wassup", n)
}

func Switch() {
	switch "Medhi" {
	case "Daniel":
		wassup("Daniel")
	case "Medhi":
		wassup("Medhi")
		fallthrough
	case "other":
		wassup("other")
		fallthrough
	default:
		fmt.Println("have you no friends?")
	}
}

type Contact struct {
	greeting string
	name     string
}

func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("Contact")
	default:
		fmt.Println("Unknown")
	}
}
