package main

import "fmt"

const mod = 1000000000

var res, n int
var tree [100009][]Edge

type Edge struct {
	to     int
	weight int
}

func main() {
	fmt.Scan(&n)

	var a, b, l int
	for i := 0; i < n-1; i++ {
		fmt.Scan(&a, &b, &l)
		tree[a] = append(tree[a], Edge{b, l})
		tree[b] = append(tree[b], Edge{a, l})
	}
	dfs(1, -1)
	fmt.Print(res)
}

func dfs(v, p int) int {
	var cnt int = 1
	for i := 0; i < len(tree[v]); i++ {
		to := tree[v][i].to
		w := tree[v][i].weight
		if to != p {
			c := dfs(to, v)
			res = (res + c*(n-c)*w) % mod
			cnt += c
		}
	}
	return cnt
}
