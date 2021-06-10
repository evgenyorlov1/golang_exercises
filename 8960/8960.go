package main

import "fmt"

func main()  {
	var n, val, sum int
	fmt.Scan(&n)

	min := 999
	max := -999
	var arr []int
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		arr = append(arr, val)
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}

	for i := 0; i < n; i++ {
		if arr[i] != min && arr[i] != max {
			sum += arr[i]
		}
	}
	fmt.Print(sum)
}