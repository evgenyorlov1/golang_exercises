package main

import "fmt"

func main()  {
	var n, val, count int
	fmt.Scan(&n)

	min := 999
	var arr []int
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		arr = append(arr, val)
		if min == val {
			count++
		}
		if val < min {
			min = val
			count = 0
		}
	}

	for i := 0; i <= count; i++ {
		fmt.Print(min, " ")
	}
	for _, v := range arr {
		if v != min {
			fmt.Print(v, " ")
		}
	}
}