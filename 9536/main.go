package main

import "fmt"

func main()  {
	var n, m, val int
	fmt.Scan(&n, &m)
	
	matrixA := [][]int{}
	matrixB := [][]int{}
	
	row := []int{}
	for i := 0; i < n; i++ {
		row = nil
		for j := 0; j < m; j++ {
			fmt.Scan(&val)
			row = append(row, val)
		}		
		matrixA = append(matrixA, row)
	}
	fmt.Scan()
	for i := 0; i < n; i++ {
		row = nil
		for j := 0; j < m; j++ {
			fmt.Scan(&val)
			row = append(row, val)
		}		
		matrixB = append(matrixB, row)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(matrixA[i][j] + matrixB[i][j], " ")
		}
		fmt.Println()
	}
}