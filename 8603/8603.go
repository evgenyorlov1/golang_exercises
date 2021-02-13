package main

import "fmt"

func main() {
	var a, b, c int
	var num string
	fmt.Scan(&num)
	
	a = int(num[0]) - 48
	b = int(num[1]) - 48
	c = int(num[2]) - 48

	fmt.Print(a+b+c, " ", a*b*c)
}