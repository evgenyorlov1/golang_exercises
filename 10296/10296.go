package main

import "fmt"

var m map[int]int

func reqursion(n int) int {
	if n == 0 {
		return 1
	}

	v2, ok2 := m[n/2]
	if !ok2 {
		v2 = reqursion(n / 2)
		m[n/2] = v2
	}

	v3, ok3 := m[n/3]
	if !ok3 {
		v3 = reqursion(n / 3)
		m[n/3] = v3
	}

	return v2 + v3
}

func main() {
	m = make(map[int]int)
	var n int
	fmt.Scan(&n)
	fmt.Print(reqursion(n))
}
