package main

import "fmt"

func main() {
	var n, val int
	fmt.Scan(&n)

	matrix := [][]int{}
	slice := []int{}

	for i := 0; i < n; i++ {
		slice = nil	
		for j := 0; j < n; j++ {
			fmt.Scan(&val)
			slice = append(slice, val)
		}
		matrix = append(matrix, slice)
	}
	
	for i := 0; i < n; i++ {
		for j := n - 1; j >= 0; j-- {
			fmt.Print(matrix[j][i], " ")	
		}
		fmt.Println()
	}
}