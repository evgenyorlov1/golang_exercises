package main

import (
	"fmt"
	"sort"
)

func main() {
	var n string
	fmt.Scan(&n)

	mp := make(map[rune]int)
	var digits []int
	for _, r := range n {
		mp[r]++
		digits = append(digits, int(r))
	}

	sort.Ints(digits)
	fmt.Print(mp[rune(digits[len(digits)-1])])

}
