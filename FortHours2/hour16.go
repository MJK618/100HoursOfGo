package FortHours2

import "fmt"

func hour16() {
	for i := 1; i <= 10000; i++ {
		fmt.Print(i, ", ")
		if i%25 == 0 {
			fmt.Println()
		}
	}

	for i := 65; i <= 90; i++ {
		fmt.Printf("%v\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}

	i := 2012
	for i <= 2021 {
		fmt.Printf("%v\n", i)
		i++

	}


	i = 2012
	for {
		if i > 2021 {
			break
		}
		fmt.Printf("%v\n", i)
		i++
	}


	i = 10
	for {
		if i > 100 {
			break
		}
		fmt.Println(i % 4)
		i++
	}


	x := 105
	if x > 100 {
		fmt.Println(">100")
	} else if x < 0 {

		fmt.Println("<0")
	} else {
		fmt.Println(">0 & <100")
	}



	switch {
	case false:
		fmt.Println("I am switch false")
	case true:
		fmt.Println("I am switch true")
	}


	favSport := "Football"
	switch {
	case favSport == "Football":
		fmt.Println("I am Football lover")
	case favSport == "Cricket":
		fmt.Println("I am Cricket lover")
	default:
		fmt.Println("I am None lover")
	}
	switch favSport  {
	case "Football":
		fmt.Println("I am Football lover")
	case "Cricket":
		fmt.Println("I am Cricket lover")
	default:
		fmt.Println("I am None lover")
	}

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
