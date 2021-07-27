package main

import "fmt"

func main() {
	var n, m int64
	fmt.Scan(&n, &m)
	if n > m {
		fmt.Print(n - m)
		return
	}
	if n < m {
		fmt.Print(m - n)
		return
	}
	fmt.Print(0)
}
