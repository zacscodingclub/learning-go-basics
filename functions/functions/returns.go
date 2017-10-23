package functions

import "fmt"

func GreetReturn(fname, lname string) string {
	return fmt.Sprint(fname, " & ", lname)
}

func MultipleReturn(fname, lname string) (string, string) {
	return fmt.Sprint(fname, " + ", lname), fmt.Sprint(lname, " * ", fname)
}
