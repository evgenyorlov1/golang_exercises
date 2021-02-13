package main

import "fmt"

func main()  {
	var n int
	fmt.Scan(&n)

	matrix := [][]int{}
	row := []int{}
	for i := 0; i < n; i++ {
		row = nil
		for j := 0; j < n; j++ {
			row = append(row, 2)
		}
		matrix = append(matrix, row)
	}
	
	j := n - 1
	for i := 0; i < n; i++ {
		matrix[i][j] = 0
		j -= 1
	}

	for i := 1; i < n; i++ {
		for j := n - 1; j > n - 1 - i; j-- {
			matrix[i][j] = 1
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
}