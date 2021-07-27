package main

import "fmt"

func main() {
	var n, count int
	fmt.Scan(&n)

	for i := 1; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				if i+j+k == n {
					count++
				}
			}
		}
	}
	fmt.Println(count)

	for i := 1; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				if i+j+k == n {
					fmt.Printf("%d%d%d\n", i, j, k)
				}
			}
		}
	}
}
