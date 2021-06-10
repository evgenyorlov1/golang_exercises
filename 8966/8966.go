package main

import "fmt"

func main() {
	var n, val int
	fmt.Scan(&n)

	var arr []int
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		arr = append(arr, val)
	}

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(arr[i], " ")
	}
}
