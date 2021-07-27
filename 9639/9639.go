package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, a, b, k int
	var keys []int
	fmt.Scan(&n)

	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		if _, ok := mp[b]; ok {
			mp[b] += a
		} else {
			keys = append(keys, b)
			mp[b] = a
		}
	}
	sort.Ints(keys)

	fmt.Scan(&k)
	for _, key := range keys {
		val, _ := mp[key]
		if k-val > 0 {
			k -= val
		} else {
			fmt.Print(key)
			break
		}
	}
}
