package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)
	fmt.Println(n[len(n)-3]-48)
}