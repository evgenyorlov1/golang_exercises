package main

import "fmt"

func main()  {
	var x, y, z float64
	fmt.Scan(&x, &y, &z)
	fmt.Printf("%.4f", x + y + z)
	fmt.Print(" ")
	fmt.Printf("%.4f", x * y * z)
}