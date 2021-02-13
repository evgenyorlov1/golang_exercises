package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&a, &b)
		fmt.Println(a + b)
	}
}