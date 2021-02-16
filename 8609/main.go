package main

import "fmt"

func recursion(n int) int {
	if n == 0 {
		return 0
	}
	return recursion(n - 1) + n
}

func main()  {
	var n int
	fmt.Scan(&n)
	fmt.Print(recursion(n))
}