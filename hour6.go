package _100HoursOfGo

import (
	"bufio"
	"fmt"
	"os"
)

func hour6() {
	//Working with maps and Dictionaries
	var mymap map[string]int
	mymap = make(map[string]int)

	mymap["India"] = 138
	mymap["China"] = 150
	mymap["Nepal"] = 10

	for key := range mymap {
		fmt.Println(key, " = ", mymap[key])
	}

	//Using Buffio and other modules for inputs and outputs
	fmt.Println("Enter a string")
	scn := bufio.NewReader(os.Stdin)
	str, err := scn.ReadString('\n')

	if err == nil {
		fmt.Println(len(str))
		str = str[:len(str)-2]

		fmt.Println(len(str))

		for i := 0; i < len(str); i++ {
			for j := i + 1; j <= len(str); j++ {
				fmt.Println("(", i, "-", j, ")", str[i:j])
			}
		}
	}
}