package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

type DoubleZero struct {
	Person
	First         string
	LicenseToKill bool
}

/*
func (r <receiver>) methodName(args type) return type {
	** magic, dragons, and things
}
*/

func (p *DoubleZero) fullName() {
	fmt.Println(p.Person.First, p.Person.Last)
}

func (p *DoubleZero) toggleLicenseToKill() {
	if p.LicenseToKill {
		p.LicenseToKill = false
	} else {
		p.LicenseToKill = true
	}
}

func main() {
	zac := DoubleZero{
		Person: Person{
			First: "Zac",
			Last:  "Baston",
			Age:   33},
		First:         "The Lifter",
		LicenseToKill: false,
	}
	zac.fullName()
	money := DoubleZero{
		Person: Person{
			First: "Miss",
			Last:  "MoneyPenny",
			Age:   29,
		},
		First:         "If Looks could kill",
		LicenseToKill: false,
	}

	m, _ := json.Marshal(money)
	var p2 Person
	bs := []byte(`{"First":"Bob", "Last":"Loblaw", "Age":21}`)
	json.Unmarshal(bs, &p2)
	fmt.Println(p2.First)
	fmt.Printf("%T \n", m)
	fmt.Println(string(m))

	// fmt.Println("Money has license to kill?", money.LicenseToKill)
	// money.toggleLicenseToKill()
	// fmt.Println("Money has license to kill?", money.LicenseToKill)
}

/*
A struct is a sequence of named elements, called fields, each of which has a name and
type.  Field names may be specified explicitly (IdentifierList) or implicitly (AnonymousField).
Within a struct, non-blank field names must be unique.

A field declared with a type but no explicit field name is an anonymous field, also called
embedded field or an embedding of the type in the struct.  An embedded type must be specified
as a type name *T, and T itself may not be a pointer type.  The unqualified type name as the
field name.

Bill Kennedy:
Go allows us the ability to declare our own types.  Struct types are declared by composing a
fixed set of unique fields together.  Each field in a struct is declared with a known type.
This could be a built-in type or another user defined type. Once we have a type declared, we
can create values from the type.  When we declare variables, the value that it represents
is anlways initialized.  The value can be initialized with a specific value or it can be
initialzied to it's zero value.  For numeric types, the zero value would be 0; for strings
it would be empty; and for bool it would be false.  In the case of a strcut, the zero value
would apply to all the different fields in the struct.  Anytime a variable is created and
initialized to it's zero value, it is idiomatic to use the keyword var.  Reserve the use of
the keyword var as a way to indicate that a variable is being set to it's zero value.  If
the variable will be initialized to something other than it's zero value, then use the short
variable declaaration operator with a struct literal.
*/
