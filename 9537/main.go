package main

import "fmt"

func main()  {
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

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(matrix[j][i], " ")
		}
		fmt.Println()
	}

}