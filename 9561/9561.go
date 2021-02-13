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

	var max int
	for i := 0; i < n; i++ {
		max = -999
		for j := 0; j < m; j++ {
			if matrix[j][i] > max {
				max = matrix[j][i]
			}
		}
		fmt.Print(max, " ")
	}
}