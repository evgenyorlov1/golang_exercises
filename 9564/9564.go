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

	var max int = -999
	row = nil
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == max {
				row = append(row, i + 1)
			}
			if matrix[i][j] > max {
				row = nil
				max = matrix[i][j]
				row = append(row, i + 1)
			}
		}
	}
	for _, v := range row {
		fmt.Print(v, " ")
	}
}