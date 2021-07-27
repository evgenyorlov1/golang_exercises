package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k, val int
	fmt.Scan(&n, &k)

	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		if _, ok := mp[val]; ok {
			mp[val]++
		} else {
			mp[val] = 1
		}
	}

	var keys []int
	for _, v := range mp {
		keys = append(keys, v)
	}

	if len(mp) < k {
		fmt.Print(k - len(mp))
		return
	}

	sort.Ints(keys)
	sum := 0
	for i := 0; i < len(mp)-k; i++ {
		sum += keys[i]
	}
	fmt.Print(sum)
}
