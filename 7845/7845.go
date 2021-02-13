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
	
	var res int
	for i := 1; i < n - 1; i++ {
		if list[i] > list[i-1] && list[i] > list[i+1] {
			res += 1
		}
	}
	fmt.Println(res)
}