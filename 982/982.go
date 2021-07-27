package main

import (
	"fmt"
)

var n, cnt int
var mt [109][109]int
var used [109]int

func dfs(v int) {
	used[v] = 1
	cnt++
	for i := 1; i <= n; i++ {
		if mt[v][i] == 1 && used[i] == 0 {
			dfs(i)
		}
	}
}

func main() {
	var m, u, v int
	fmt.Scan(&n, &m)

	for i := 0; i < m; i++ {
		fmt.Scan(&u, &v)
		mt[v][u] = 1
		mt[u][v] = 1
	}

	dfs(1)
	if cnt == n {
		fmt.Print("YES")
		return
	}
	fmt.Print("NO")
}
