package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a >= -9 && a <= 9 {
		fmt.Print("Ok")
	} else {
		fmt.Print("No")
	}
}