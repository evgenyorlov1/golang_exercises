package main

import "fmt"

func main() {
	var n, val int
	fmt.Scan(&n)

	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		if _, ok := mp[val]; ok {
			mp[val]++
		} else {
			mp[val] = 1
		}

		if _, ok := mp[val-1]; ok {
			mp[val-1]++
		} else {
			mp[val-1] = 1
		}

		if _, ok := mp[val+1]; ok {
			mp[val+1]++
		} else {
			mp[val+1] = 1
		}
	}

	max := -999
	for _, v := range mp {
		if max < v {
			max = v
		}
	}
	fmt.Print(max)
}
