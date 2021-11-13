package FortHours1

import (
	"fmt"
)

//Slicing
func hour4 () {
	fmt.Println("Enter a number")
	var n int
	fmt.Scanln(&n)

	var slice = make([]int, 0)
	appendOddsTillNToSlice(n, &slice)

	var len int
	len = display(slice)
	fmt.Println("Size of slice is ", len)

	fmt.Println("Enter a number")
	var m int
	fmt.Scanln(&m)

	mutiplyBy2Slice(m, &slice)
	display(slice)

	var remove5th = append(slice[:5], slice[6:]...)
	display(remove5th)
}
//Functions
func appendOddsTillNToSlice(n int, slice *[]int) {
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			*slice = append(*slice, i)
		}
	}
}

func display(slice []int) int {
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i], " -> ", i, "| ")
	}
	fmt.Println(".")
	return len(slice)
}

func mutiplyBy2Slice(m int, slice *[]int) {
	for i := 0; i < len(*slice); i++ {
		if (*slice)[i] >= m {
			(*slice)[i] = (*slice)[i] * 2
		}
	}
}