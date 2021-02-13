package main

import "fmt"

func main()  {
	var m, n, x, y int
	fmt.Scan(&m, &n, &x, &y)

	if x - 1 >= 1 {
		fmt.Println(x - 1, " ", y)
	}
	if x + 1 <= m {
		fmt.Println(x + 1, " ", y)
	}
	if y - 1 >= 1 {
		fmt.Println(x, " ", y - 1)
	}
	if y + 1 <= n {
		fmt.Println(x, " ", y + 1)
	}
}