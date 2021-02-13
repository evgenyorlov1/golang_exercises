package main

import "fmt"

func main() {
	var c string
	var n, s, r, res int
	fmt.Scan(&c)
	fmt.Scan(&n)

	res = ((2 + 3 * (n - 1)) * n) / 2
	fmt.Println(res)
	s = n -1
	r = 1
	for i := 0; i < n; i++ {
		for j := 0; j < s; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < r ; k++ {
			fmt.Print(c)
		}
		fmt.Println()
		r += 2
		s -= 1
	}
}