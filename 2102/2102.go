package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	if n < k && m < k {
		fmt.Println(1)
		return
	}

	a := n / k
	if n%k >= 1 {
		a++
	}

	b := m / k
	if m%k >= 1 {
		b++
	}

	fmt.Println(a * b)
}
