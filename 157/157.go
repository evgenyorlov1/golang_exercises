package main

import "fmt"

func main() {
	var n, res int
	fmt.Scanln(&n)
	res = (n-1)*(n-2)
	fmt.Println(res)
}