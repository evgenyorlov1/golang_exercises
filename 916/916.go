package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a > b {
		a, b = b, a
	}
	if c > d {
		c, d = d, c
	}

	mp := make(map[int]int)
	for i := a; i <= b; i++ {
		for j := c; j <= d; j++ {
			mp[i*j]++
		}
	}

	fmt.Print(len(mp))
}
