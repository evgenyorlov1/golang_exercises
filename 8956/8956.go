package main

import "fmt"

func main()  {
	var n, val int
	fmt.Scan(&n)

	var arr []int
	count := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		if val < 0 {
			count++
			arr = append(arr, val)
		}
	}

	if count == 0 {
		fmt.Println("NO")
		return
	}
	
	fmt.Println(count)
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(arr[i], " ")
	}
}