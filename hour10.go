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
}

func foo() {
	fmt.Println("again:", y)
	fmt.Println(zeta)
}
