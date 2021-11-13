package FortHours1

import "fmt"

func main(){
	fmt.Println("Enter a number")
	var n int
	fmt.Scanln(&n)

	hour5(n)
}

func hour5(n int){
	if n == 0{
		return;
	}

	fmt.Println(n)
	hour5(n - 1)
	fmt.Println(n)
}