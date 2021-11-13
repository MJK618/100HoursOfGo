package FortHours1

import "fmt"

func hour3(){

	//1D Array
	fmt.Println("Enter the size of array")
	var num int
	fmt.Scanln(&num)
	var slice = make([]int, num)
	for i := 0; i < len(slice); i++ {
		fmt.Println(i, "th number please?")
		fmt.Scanln(&slice[i])
	}
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i], "\t")
	}

	//2D Array
	var a [3][4]int
	a[0][0] = 11
	a[0][1] = 12
	a[0][2] = 13
	a[0][3] = 14
	a[1][0] = 21
	a[1][1] = 22
	a[1][2] = 23
	a[1][3] = 24
	a[2][0] = 31
	a[2][1] = 32
	a[2][2] = 33
	a[2][3] = 34
	for i := 0; i < len(a); i++ {
		for j:= 0; j < len(a[0]); j++ {
			fmt.Print(a[i][j], "\t")
		}
		fmt.Println()
	}
}