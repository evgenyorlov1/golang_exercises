package main

import "fmt"

func main() {
	var n, a, b, c, d int
	fmt.Scan(&n, &a, &b, &c, &d)

	var arr []int
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}

	for a < b {
		arr[a-1], arr[b-1] = arr[b-1], arr[a-1]
		a++
		b--
	}

	for c < d {
		arr[c-1], arr[d-1] = arr[d-1], arr[c-1]
		c++
		d--
	}

	for _, v := range arr {
		fmt.Print(v, " ")
	}
}
