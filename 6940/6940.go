package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	m := make(map[string]int)
	var word string
	for i := 0; i < n; i++ {
		fmt.Scan(&word)
		_, ok := m[word]
		if ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}

	max := -999
	for _, v := range m {
		if v > max {
			max = v
		}
	}

	var arr []string
	for k, v := range m {
		if v == max {
			arr = append(arr, k)
		}
	}
	sort.Strings(arr)
	fmt.Print(arr[len(arr)-1], " ", max)
}
