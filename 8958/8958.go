package main

import "fmt"

func main()  {
	var n, val int
	fmt.Scan(&n)

	var arr []int
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		if i % 2 == 1 {
			arr = append(arr, val)
		}
	}

	if len(arr) == 0 {
		fmt.Print("NO")
		return
	}

	fmt.Println(len(arr))
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
}