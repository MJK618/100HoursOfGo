package _100HoursOfCode

import (
	"fmt"
	"strconv"
)

func hour2() {
	fmt.Println("Hour 2 Started")

	//Take Input and Print Outputs
	fmt.Println("Please tell your name")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Hello " + name)

	//Condtionals
	fmt.Println("Enter a number")
	var num int
	fmt.Scanln(&num)
	if num % 2 == 0 {
		fmt.Println(strconv.Itoa(num) + " is an even number")
	} else {
		fmt.Println(strconv.Itoa(num) + " is an odd number")
	}
}
