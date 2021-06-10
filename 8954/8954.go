package main;

import "fmt"

func main()  {
	var n, val int
	
	fmt.Scanln(&n)
	
	var arr []int
	for i := 1; i <= n; i++ {
		fmt.Scanln(&val)
		arr = append(arr, val)
	}

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(arr[i], " ")
	}
}
