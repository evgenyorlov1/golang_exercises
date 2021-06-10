package main;

import "fmt"

func main()  {
	var n, val int
	
	fmt.Scanln(&n)
	
	var arr []int
	for i := 1; i <= n; i++ {
		fmt.Scan(&val)
		arr = append(arr, val)
	}

	for _, v := range arr {
		fmt.Println(v)
	}
}
