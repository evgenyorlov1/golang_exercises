package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scanf("%d", &n)

	if n > 0 {
		for i := 1; i <= n; i++ {
			sum += i
		}
	} else {
		for i := 1; i >= n; i-- {
			sum += i
		}
	}
	fmt.Println(sum)
}