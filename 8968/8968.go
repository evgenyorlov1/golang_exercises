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
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	max = max / 2
	min = min / 2

	for _, v := range arr {
		if v > 0 {
			v = v - max
		}
		if v < 0 {
			v = v - min
		}
		fmt.Print(v, " ")
	}
}
