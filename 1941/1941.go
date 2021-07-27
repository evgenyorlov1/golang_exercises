package main

import "fmt"

var used [100009]int
var enter [100009]int
var leave [100009]int
var time int

func isParent(a, b int) int {
	if enter[a] < enter[b] && leave[a] > leave[b] {
		return 1
	}
	return 0
}

func dfs(v int, gr [][]int) {
	used[v] = 1
	time++
	enter[v] = time
	for i := 0; i < len(gr[v]); i++ {
		to := gr[v][i]
		if used[to] == 0 {
			dfs(to, gr)
		}
	}
	time++
	leave[v] = time
}

func main() {
	gr := make([][]int, 100009)
	for i := 0; i < 100009; i++ {
		gr[i] = make([]int, 100009)
	}

	var n, val, root int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&val)
		if val == 0 {
			root = i
		} else {
			gr[val] = append(gr[val], i)
		}
	}
	dfs(root, gr)
	var m, a, b int
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		fmt.Scan(&a, &b)
		fmt.Println(isParent(a, b))
	}
}
