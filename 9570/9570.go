package main

import "fmt"

func main() {
	var n, m, val int
	fmt.Scan(&n, &m)
	slice := []int{}

	for i := 0; i < n; i++ {
		slice = nil
		for j := 0; j < m; j++ {
			fmt.Scan(&val)
			slice = append(slice, val)
		}
		for j := m - 1; j >= 0; j-- {
			fmt.Print(slice[j], " ")	
		}
		fmt.Println()
	}
}