package main

import "fmt"

func main() {
	var n, val int
	fmt.Scan(&n)

	var arr []int
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		_, ok := m[val]
		if ok {
			arr = append(arr, val)
			m[val]++
		} else {
			m[val] = 1
		}
	}

	if len(arr) == 0 {
		fmt.Print("NO")
		return
	}

	var result []int
	for i := len(arr) - 1; i >= 0; i-- {
		_, ok := m[arr[i]]
		if ok {
			result = append(result, arr[i])
		}
		delete(m, arr[i])
	}

	for i := len(result) - 1; i >= 0; i-- {
		fmt.Print(result[i], " ")
	}
}
