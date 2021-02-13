package main

import "fmt"

func main() {
	var n, val int
	fmt.Scan(&n)

	var list = []int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		list = append(list, val)
	}
	for i, v := range list {
		if i % 2 == 0 {
			fmt.Print(v, " ")
		}
	}
}