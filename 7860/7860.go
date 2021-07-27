package main

import "fmt"

func main() {
	var n int
	var word string
	fmt.Scan(&n)

	DOMjudge := make(map[string]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&word)
		_, ok := DOMjudge[word]
		if ok {
			DOMjudge[word]++
		} else {
			DOMjudge[word] = 1
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&word)
		v, ok := DOMjudge[word]
		if ok && v > 0 {
			DOMjudge[word]--
			count++
		}
	}

	fmt.Print(count)
}
