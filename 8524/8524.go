package main

import "fmt"

func main() {
	var n, val int
	fmt.Scan(&n)

	res := 0
	for i := 0; i < n*n; i++ {
		if fmt.Scan(&val); val > 0 {
			res += val
		}
	}
	fmt.Println(res)
}
