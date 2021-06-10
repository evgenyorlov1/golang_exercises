package main

import "fmt"

func main()  {
	var n, val, index int
	fmt.Scan(&n)

	var arr []int
	min := 999
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		arr = append(arr, val)
		if val < min {
			min = val
			index = i
		}
	}

	arr[index] = arr[0]
	arr[0] = min
	for _, v := range arr {
		fmt.Print(v, " ")
	}
}