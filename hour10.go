package _100HoursOfGo

import (
	"fmt"
)

// DECLARE the variable "y"
// ASSIGN the value 43
// declare & assign = initialization
var y = 43

// DECLARE there is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
// false for booleans, 0 for integers, 0.0 for floats, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels, and maps.
var z int
var zeta = 21

var a int
var b string = "James Bond"


func hour10() {
	// short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	x := 42
	fmt.Println(x)
	fmt.Println(y)
	foo()
	fmt.Println(z)

	xeta := 42 + 7
	yeta := "James Bond"
	fmt.Println(xeta)
	fmt.Println(yeta)
	xeta = 50
	fmt.Println(xeta)
	fmt.Println(zeta)


	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\t%T\n", a, b)

	s := fmt.Sprint(a, " something more ", b)
	fmt.Println(s)
	s2 := fmt.Sprintf("%v\t%T\t%T\n", "to pass in", a, b)
	fmt.Println(s2)
}

func foo() {
	fmt.Println("again:", y)
	fmt.Println(zeta)
}
