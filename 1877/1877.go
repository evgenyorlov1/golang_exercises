package main

import "fmt"

func main() {
	var a, b int64
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	if a % 2 == 0 || b % 2 != 0 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}