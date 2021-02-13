package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if 45*(12-a) <= 255{
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}