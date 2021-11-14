package FortHours2

import "fmt"

func hour18() {
	//Arrays
	a := [5]int{1, 4, 2, 4, 7}
	fmt.Println(a)

	for i, v := range a {
		fmt.Println(i,v)
	}
	fmt.Printf("%T", a)

	//Slices
	b := []string{"one", "two", "three", "four", "five", "six", "seven", "eigth", "nine", "ten"}
	fmt.Println(b)

	for i, v := range b {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", b)

	c:= []string{"one", "two", "three", "four", "five", "six", "seven", "eigth", "nine", "ten"}
	fmt.Println(b)

	fmt.Println(c[:5])
	fmt.Println(c[5:])
	fmt.Println(c[2:7])
	fmt.Println(c[1:6])

	d := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	d = append(d, 52)
	fmt.Println(d)
	d = append(d, 53, 54, 55)
	fmt.Println(d)

	y := []int{56, 57, 58, 59, 60}
	d = append(d, y...)
	fmt.Println(d)


	e := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(e)
	f := append(e[0:3], e[6:]...)
	fmt.Println(f)

	g := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`,
		` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`,
		` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`,
		` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`,
		` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`,
		` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`,
		` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`,
		` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`,
		` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	fmt.Println(g)

	fmt.Println(len(g))
	fmt.Println(cap(g))

	for i := 0; i < len(g); i++ {
		fmt.Println(i, g[i])
	}
	h := []string{"James", "Bond", "Shaken, not stirred"}
	k := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(h)
	fmt.Println(k)

	l:= [][]string{h, k}
	fmt.Println(l)

	for i, m:= range l{
		fmt.Println("record: ", i)
		for j, val := range m{
			fmt.Printf("\t index position: %v value: %v \n", j, val)
		}
	}
	n := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	fmt.Println(n)

	for k, v := range n {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
	o := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	fmt.Println(o)
	n["biden_joe"] = []string{`Democratic pig`, `China`, `Russia`}

	for k, v := range o {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

	//Adding new record
	o["biden_joe"] = []string{`Democratic pig`, `China`, `Russia`}
	fmt.Println("\n\nAfter Adding the new record")
	for k, v := range o {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

	p:= map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	fmt.Println(p)

	for k, v := range p{
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

	//Deleting a record
	delete(p,"no_dr")
	fmt.Println("\n\nAfter Deleting record")
	for k, v := range p{
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}