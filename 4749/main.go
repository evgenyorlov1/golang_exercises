package main

import "fmt"

func main() {
	var n, m, val, res int
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
	fmt.Scanln()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&val)
			res += matrix[i][j] * val
		}
	}
	fmt.Print(res)
}