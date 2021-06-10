package main

import "fmt"

func main()  {
	var n, val int
	fmt.Scan(&n)

	max := -999
	min := 999

	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	fmt.Print(max - min)
}