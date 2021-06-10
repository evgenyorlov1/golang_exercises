package main

import "fmt"

func main()  {
	var n, val, index int
	fmt.Scan(&n)

	max := -999
	var arr []int
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		arr = append(arr, val)
		if max <= val {
			max = val
			index = i
		}
	}

	arr[index] = arr[n-1]
	arr[n-1] = max
	for _, v := range arr {
		fmt.Print(v, " ")
	}
}