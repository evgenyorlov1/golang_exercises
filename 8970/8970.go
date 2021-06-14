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
	fmt.Println(n / 2)
	for i := 0; i < n/2; i++ {
		fmt.Print(arr[i+n/2], " ", arr[i], " ")
	}
}
