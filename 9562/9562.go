package main

import "fmt"

func main(){
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

	var r1, c1, r2, c2 int	
	fmt.Scan(&r1, &c1, &r2, &c2)
	var sum int
	for i := r1 - 1; i <= r2 - 1; i++ {
		for j := c1 - 1; j <= c2 - 1; j++ {
			sum += matrix[i][j]
		}
	}
	fmt.Println(sum)
}