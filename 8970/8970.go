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
	even := n/2 + n%2
	for i := 0; i < even; i++ {
		if i+even > len(arr)-1 {
			fmt.Print(arr[i])
		} else {
			fmt.Print(arr[i], " ", arr[i+even], " ")
		}
	}
}
