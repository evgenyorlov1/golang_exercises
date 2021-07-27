package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n)

	for i := 0; i < n/2; i++ {
		fmt.Scan(&a, &b)
		fmt.Print(b, " ", a, " ")
	}

	if n%2 == 1 {
		fmt.Scan(&a)
		fmt.Print(a)
	}
}
