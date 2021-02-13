package main

import "fmt"

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func main()  {
	var n int
	var val, res float64
	res = 111.0
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		res = min(res, val)
	}
	fmt.Printf("%.2f", 2*res)
}