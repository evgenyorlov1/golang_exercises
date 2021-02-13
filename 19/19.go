package main

import "fmt"

func main() {
	var count int
	var n string
	fmt.Scan(&n)

	if len(n) % 2 == 1 {
		count += 1
	}
	w := len(n) / 2
	for i := 0; i < w; i++ {
		if n[i] == n[len(n) - i - 1] {
			count += 1
		}
	}
	fmt.Print(count)
}