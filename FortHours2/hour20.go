package FortHours2

import "fmt"

type person struct {
	first_name   string
	last_name    string
	fav_icecream []string
}
type vehicle struct {
	color string
	doors string
}

type truck struct {
	vehicle
	fourWheels bool
}

type sedan struct {
	vehicle
	luxury bool
}
func hour20() {
	p1 := person{
		first_name: "Jatin",
		last_name:  "Kamboj",
		fav_icecream: []string{"Vanilla",
			"Choclate"},
	}
	p2 := person{
		first_name: "Todd",
		last_name:  "Mcloedd",
		fav_icecream: []string{"Strawberry",
			"Brownie"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	m := map[string]person{p1.last_name: p1,
		p2.last_name: p2}
	fmt.Println(m)

	v1 := truck{
		vehicle: vehicle{
			color: "Red",
			doors: "2",
		}, fourWheels: true,
	}
	v2 := sedan{
		vehicle: vehicle{
			color: "white",
			doors: "4",
		}, luxury: true,
	}

	fmt.Println(v1)
	fmt.Println(v2)

	x:= struct {
		color string
		doors int
	}{
		color: "Red",
		doors: 2,
	}

	fmt.Println(x)
}

