package main

import "fmt"

const max int = 10000000

var arr [max]int
var m map[int]int

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func f(n int) int {
	if n < max {
		return arr[n]
	}
	if m[n] > 0 {
		return m[n]
	}
	if n%3 == 0 {
		m[n] = f(n/3) + 5
	} else {
		m[n] = f(n-4) + 2
	}
	return m[n]
}

func main() {
	arr[0] = -max
	arr[1] = 0
	arr[2] = -max
	arr[3] = 5
	arr[4] = -max
	arr[5] = 2
	for i := 6; i < 110; i++ {
		if i%3 == 0 {
			arr[i] = min(arr[i/3]+5, arr[i-4]+2)
			fmt.Print(arr[i/3]+5, " ", arr[i-4]+2, " ", min(arr[i/3]+5, arr[i-4]+2), " ")
		} else {
			arr[i] = m[i-4] + 2
		}
		fmt.Print(arr[i], "\n")
	}
	var n int
	fmt.Scan(&n)

	res := f(n)
	if res < 0 {
		res = 0
	}
	fmt.Print(res)
}
