package main

import "fmt"

func main() {
	var n, res, y, x int
	fmt.Scan(&n)

	var pos string
	for i:=0; i<n; i++ {
		res = 0
		fmt.Scan(&pos)
		x = int(pos[0]) - 96
		y = int(pos[1]) - 48

		if y+2 <= 8 && x-1 >= 1 {
			res += 1
		}
		if y+2 <= 8 && x+1 <= 8 {
			res += 1
		}
		if y-2 >= 1 && x-1 >= 1 {
			res += 1
		}
		if y-2 >= 1 && x+1 <= 8 {
			res += 1
		}
		if y+1 <= 8 && x+2 <= 8 {
			res += 1
		}
		if y-1 >= 1 && x+2 <= 8 {
			res += 1
		}
		if y+1 <= 8 && x-2 >= 1 {
			res += 1
		}
		if y-1 >= 1 && x-2 >= 1 {
			res += 1
		}
		fmt.Println(res)
	}
}