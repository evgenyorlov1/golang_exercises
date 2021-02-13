package main

import "fmt"

func appendUnique(slice []int, val int) []int {
	for _, v := range slice {
		if v == val {
			return slice
		}
	}	
	slice = append(slice, val)
	return slice
}

func main() {
	var n, m, val int
	fmt.Scan(&n, &m)
	
	var min int = 999
	rows := []int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&val)
			if val == min {
				rows = appendUnique(rows, i + 1)
			}
			if val < min {
				rows = nil
				min = val
				rows = appendUnique(rows, i + 1)
			}
		}
	}

	for _, v := range rows {
		fmt.Print(v, " ")
	}
}