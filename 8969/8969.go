package main

import "fmt"

func main() {
	var n, val int
	fmt.Scan(&n)

	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		if i == 0 {
			fmt.Print(val, " ")
			sum = val
		} else {
			fmt.Print(val-sum, " ")
			sum += val - sum
		}
	}
}
