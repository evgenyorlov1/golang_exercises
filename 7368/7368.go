package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m, val, sum, count int
	var a, b int
	var res float64
	fmt.Scan(&n, &m)

	list := []int{}
	for i := 0; i < m; i++ {
		sum  = 0
		list = nil
		count = 0
		
		for j := 0; j < n; j ++ {
			fmt.Scan(&val)
			list = append(list, val)
		}

		sort.Ints(list)
		a = list[0]
		b = list[len(list) - 1]
		for k := 0; k < n; k++ {
			if list[k] != a && list[k] != b {
				sum += list[k]
				count += 1
			}
		}
		
		res = float64(sum) / float64(count)
		fmt.Printf("%.2f", res)
		fmt.Print(" ")
	}
}