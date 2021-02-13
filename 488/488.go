package main

import "fmt"

func main() {
	var n, step int
	fmt.Scan(&n)

	for row := 1; row <= n; row++ {
		if row % 2 == 1 {
			for col := 1; col <= n; col++ {
				fmt.Print(col + step, " ")
			}
		} else {
			for col := n; col >= 1; col-- {
				fmt.Print(col + step, " ")
			}
		}
		fmt.Println()
		step += n
	}
}