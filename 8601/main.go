package main

import "fmt"

func main() {
	var num string
	fmt.Scan(&num)
	fmt.Printf("%c", num[1])
	fmt.Printf("%c", num[0])
}