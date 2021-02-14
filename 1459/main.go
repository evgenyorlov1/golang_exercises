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
	sort.Ints(slice)
	fmt.Print(slice[0], " ", slice[1])
}