package main

import "fmt"

func main() {
	var n, m, val int
	fmt.Scan(&n, &m)
	
	matrix := [][]int{}
	row := []int{}
	for i := 0; i < n; i++ {
		row = nil
		for j := 0; j < m; j++ {
			fmt.Scan(&val)
			row = append(row, val)
		}
		matrix = append(matrix, row)
	}

	maxs := []int{}
	var max, min int
	for i := 0; i < n; i++ {
		max = -999
		for j := 0; j < m; j++ {
			if max < matrix[i][j] {
				max = matrix[i][j]
			}
		}
		maxs = append(maxs, max)
	}
	
	min = 999
	for _, v := range maxs {
		if v < min {
			min = v
		}
	}
	
	fmt.Println(min)
}