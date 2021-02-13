package main

import "fmt"

func main() {
	var n, count, val int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		count = 0
		for j := 0; j < n; j++ {
			fmt.Scan(&val)
			count += val
		}
		fmt.Println(count)
	}
}