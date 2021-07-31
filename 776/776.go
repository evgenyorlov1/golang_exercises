package main

import "fmt"

var mt [10001][]int
var used [10001]int

func dfs(v int) {
	used[v] = 1
	for i := 0; i < len(mt[v]); i++ {
		to := mt[v][i]
		if used[to] == 0 {
			dfs(to)
		}
	}
}

func main() {
	var n, m, a, b int
	fmt.Scan(&n, &m)

	for i := 0; i < m; i++ {
		fmt.Scan(&a, &b)
		mt[a] = append(mt[a], b)
		mt[b] = append(mt[b], a)
	}

	c := 0
	for i := 1; i <= n; i++ {
		if used[i] == 0 {
			dfs(i)
			c++
		}
	}
	fmt.Print(c - 1)
}
