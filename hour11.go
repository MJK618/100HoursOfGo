package _100HoursOfGo

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var c bool = true
type MJK int

var elon MJK
type MJK int

var haha MJK

var jaja int
func hour11() {
	xeta := 42
	yeta := "James Bond"
	zeta := true
	fmt.Printf("%v, %T, %v, %T,%v, %T", xeta, xeta, yeta, yeta, zeta, zeta)
	s := fmt.Sprintf("%v, %T, %v, %T,%v, %T", x, x, y, y, c, c)
	fmt.Println(s)

	fmt.Printf("%v, %T\n", elon, elon)
	elon = 42
	fmt.Printf("%v, %T", elon, elon)


	fmt.Printf("%v, %T\n", haha, haha)
	haha = 42
	fmt.Printf("%v, %T\n", haha, haha)

	jaja = int(x)
	fmt.Printf("%v, %T", jaja, jaja)
}
