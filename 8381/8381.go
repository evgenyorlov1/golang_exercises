package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n > 10 {
		n = 10
	}

	res := 0
	if n >= 1 {
		res += 9
	}
	if n >= 2 {
		res += 9 * 9
	}
	if n >= 3 {
		res += 9 * 9 * 8
	}
	if n >= 4 {
		res += 9 * 9 * 8 * 7
	}
	if n >= 5 {
		res += 9 * 9 * 8 * 7 * 6
	}
	if n >= 6 {
		res += 9 * 9 * 8 * 7 * 6 * 5
	}
	if n >= 7 {
		res += 9 * 9 * 8 * 7 * 6 * 5 * 4
	}
	if n >= 8 {
		res += 9 * 9 * 8 * 7 * 6 * 5 * 4 * 3
	}
	if n >= 9 {
		res += 9 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2
	}
	if n >= 10 {
		res += 9 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	}
	fmt.Print(res)
}
