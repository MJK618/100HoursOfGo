package FortHours2

import "fmt"

var xolo int = 10
type person struct {
	first string
	last  string
	age   int
}
type SQUARE struct {
	side float64
}

func (sq SQUARE) area() float64 {
	return sq.side * sq.side
}

type CIRCLE struct {
	radius float64
}

func (cr CIRCLE) area() float64 {
	return 3.14 * cr.radius * cr.radius
}

type SHAPE interface {
	area() float64
}

func info(sh SHAPE) float64 {
	return sh.area()
}

func hour23() {
	defer foo2()
	fmt.Println("I am just the defer")
	fmt.Println(foo())
	fmt.Println(bar())
	fmt.Println(foo1(1,2,3,4,5,6,7,8,9,10))
	fmt.Println(bar1([] int {1,2,3,4,5,6,7,8,9,}))


	p1 := person{
		"John",
		"Lenon",
		33,
	}
	fmt.Println(p1)
	p1.speak()

	square1 := SQUARE{
		side: 4.0,
	}
	circle1 := CIRCLE{
		radius: 7.0,
	}

	fmt.Println("Area of Square is", info(square1))
	fmt.Println("Area of Circle is", info(circle1))


	func() {
		fmt.Println("Hello, playground I am in anonymous function")
	}()

	anony:= func() {
		fmt.Println("Hello, playground I am anony, a function assigned to a variable")
	}

	anony()

	x:= a()
	x()



	g := func(x int) int {
		return x + 10
	}

	fmt.Println(foo5(g, 5))

	fmt.Println(xolo)
	something(8)
	fmt.Println(xolo)
	something(18)
	fmt.Println(xolo)
}
func something(y int) {
	xolo := y
	fmt.Println(xolo)
}

func foo() int {
	return 4
}

func bar() (int, string) {
	return 12, "A word"
}

func foo1(x ...int) int {
	sum := 0
	for _,v := range x {
		sum += v
	}
	return sum
}

func bar2(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}


func foo2() {
	defer func(){
		fmt.Println("I am the foo defer")
	}()
	fmt.Println("I am the foo")

}

func (p person) speak() {
	fmt.Printf("I am %v %v, and I am %v years old", p.first, p.last, p.age)
}
func a() func() {
	defer fmt.Println("a ended")
	fmt.Println("a started")
	return func(){
		fmt.Println("Printing via anonymous function inside a")
	}
}

func foo5(f func(x int) int, i int) int {
	n := f(i)
	n++
	return n
}

