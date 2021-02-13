package main

import "fmt"

func main()  {
	var n, val int
	fmt.Scan(&n)
	
	slice := []int{}
	var min, max int = 111, -111
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		slice = append(slice, val)
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	for _, v := range slice {
		if v == max {
			fmt.Print(min, " ")
			continue
		}
		if v == min {
			fmt.Print(max, " ")
			continue
		}
		fmt.Print(v, " ")
	}
}