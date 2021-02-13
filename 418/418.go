package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	res := math.Pow(math.Sqrt(a) + math.Sqrt(b) + math.Sqrt(c), 2)
	fmt.Printf("%.8f", res)
}