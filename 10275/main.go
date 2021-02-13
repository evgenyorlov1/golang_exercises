package main

import "fmt"

func main()  {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a + b == c {
		fmt.Print("Yes")
		return
	}
	if a + c == b {
		fmt.Print("Yes")
		return
	}
	if c + b == a {
		fmt.Print("Yes")
		return
	}
	fmt.Print("No")
}