package main

import "fmt"

func main() {
	var n, val int
	fmt.Scan(&n)

	var arr []int
	var m map[int]int
	m = make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&val)

		_, ok := m[val]
		if ok {
			continue
		} else {
			arr = append(arr, val)
			m[val] = 1
		}
	}

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
}
