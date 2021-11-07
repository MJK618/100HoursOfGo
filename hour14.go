package _100HoursOfGo

import "fmt"
const (
	n = 10
	o = "an untyped string"
	p = 45.20
)

const (
	q int = 100
	r string = "a typed string"
	s float64 = 4.2069
)
const (
	currentYear = iota + 2021
	nextYear
	nextNextYear
	nextNextNextYear
)
func main() {
	a := 10
	b := 20
	g := a==b
	h := a <= b
	i:= a >= b
	j:= a != b
	k:= a < b
	l:= a > b

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

	m := 618

	fmt.Printf("%d\n%b\n%x",m,m,m)



	fmt.Println(n)
	fmt.Println(o)
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(s)



	t := 420
	fmt.Printf("%d\n%b\n%x\n",t,t,t)

	u := t << 1
	fmt.Printf("%d\n%b\n%x",u,u,u)


	v := `"So this is 
	
	a
	
	raw string literal"`

	fmt.Println(v)


	fmt.Println(currentYear)
	fmt.Println(nextYear)
	fmt.Println(nextNextYear)
	fmt.Println(nextNextNextYear)
}
