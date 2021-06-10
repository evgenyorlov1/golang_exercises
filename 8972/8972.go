package main

import "fmt"

func main() {
	var n, val int
	fmt.Scan(&n)

	var m = make(map[int]int)
	var arr []int
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		_, ok := m[val]
		if ok {
			m[val]++
		} else {
			arr = append(arr, val)
			m[val] = 1
		}
	}

	if len(m) == n {
		fmt.Println("NO")
		return
	}

	for _, v := range arr {
		count, ok := m[v]
		if ok && count > 1 {
			fmt.Print(v, " ")
		}
	}
}
