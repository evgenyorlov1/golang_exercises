package main

import (
	"fmt"
	"math"
)

func main()  {
	var val float64
	var n int
	fmt.Scan(&n)
	min := 999.0
	max := -999.0

	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		min = math.Min(min, val)
		max = math.Max(max, val)
	}
	fmt.Printf("%.0f", min + max)
}