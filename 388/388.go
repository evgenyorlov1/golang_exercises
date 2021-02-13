package main

import "fmt"

func main() {
	var n, i int
	fmt.Scan(&n)
	for i = 0; n > 1; i++ {
		if n % 2 == 1 {
			n += 1
		} else {
			n /= 2
		}
	}
	fmt.Println(i)
}