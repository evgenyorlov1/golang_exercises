package main

import (
	"fmt"
	"sort"
)

func main()  {
	var n, val int
	fmt.Scan(&n)

	slice := []int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		slice = append(slice, val)
	}

	odd := []int{}
	even := []int{}
	for _, v := range slice {
		if v % 2 == 0 {
			even = append(even, v)
			continue
		}
		odd = append(odd, v)
	}

	sort.Ints(odd)
	for _, v := range odd {
		fmt.Print(v, " ")
	}
	sort.Ints(even)
	for i := len(even) - 1; i >= 0; i-- {
		fmt.Print(even[i], " ")
	}
}