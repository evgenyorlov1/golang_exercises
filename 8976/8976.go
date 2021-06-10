package main

import "fmt"

func main() {
	var n, val int
	fmt.Scan(&n)

	var arr []int
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		arr = append(arr, val)
		_, ok := m[val]
		if ok {
			m[val]++
		} else {
			m[val] = 1
		}
	}

	max := -1
	for _, v := range m {
		if max < v {
			max = v
		}
	}

	fmt.Println(max)
	for _, v := range arr {
		if m[v] == max {
			fmt.Print(v, " ")
			delete(m, v)
		}
	}

}
