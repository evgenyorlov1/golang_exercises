package main

import "fmt"

var m [111][111]int

func dfs(x, y int) {
	if m[x][y] == 0 {
		return
	}
	if (x < 0) || (y < 0) || (x > 111) || (y > 111) {
		return
	}
	m[x][y] = 0

	dfs(x+1, y)
	dfs(x+1, y+1)
	dfs(x-1, y)
	dfs(x-1, y+1)
	dfs(x, y+1)
	dfs(x, y-1)
	dfs(x-1, y-1)
	dfs(x+1, y-1)
}

func main() {
	var n, x, y int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		m[x][y] = 1
	}

	c := 0
	for x := 0; x < 111; x++ {
		for y := 0; y < 111; y++ {
			if m[x][y] == 1 {
				dfs(x, y)
				c++
			}
		}
	}
	fmt.Println(c)
}
