package main

import "fmt"

func main()  {
	var n, val, count int
	fmt.Scan(&n)

	max := -999
	var arr []int
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		arr = append(arr, val)
		if max == val {
			count++
		}
		if val > max {
			max = val
			count = 0
		}
	}

	for _, v := range arr {
		if v != max {
			fmt.Print(v, " ")
		}
	}
	for i := 0; i <= count; i++ {
		fmt.Print(max, " ")
	}
}