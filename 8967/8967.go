package main

import "fmt"

func main() {
	var n, val int
	fmt.Scan(&n)

	var arr []int
	min := 999
	max := -999
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		arr = append(arr, val)
		if min > val {
			min = val
		}
		if max < val {
			max = val
		}
	}
	max = -(max - 2*min)
	for _, v := range arr {
		fmt.Print(v-min+max, " ")
	}
}
