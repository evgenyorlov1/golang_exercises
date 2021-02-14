package main

import "fmt"

func main()  {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	fmt.Printf("%.4f", a + b)
	fmt.Print(" ")
	fmt.Printf("%.4f", a + c)
	fmt.Print(" ")
	fmt.Printf("%.4f", c + b)
}