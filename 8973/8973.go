package main

import "fmt"

func main() {
	var n, val int
	fmt.Scan(&n)

	var m = make(map[int]int)
	var arr []int
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		m[val]++
		arr = append(arr, val)
	}

	for _, v := range arr {
		count, ok := m[v]
		if ok && count == 1 {
			fmt.Print(v, " ")
		}
		m[v]--
	}
}
