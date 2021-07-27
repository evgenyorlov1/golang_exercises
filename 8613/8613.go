package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	if x >= 13 {
		fmt.Print(3*x*x*x + 4*x*x + 5*x + 6)
		return
	}
	fmt.Print(3*x*x*x - 2*x*x - 3*x - 4)
}
