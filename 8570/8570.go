package main

import "fmt"

func main() {
	var w string

	for {
		n, err := fmt.Scanf("%s", &w)
		if n == 0 || err != nil {
			break
		}
		fmt.Print(len(w), " ")
	}
}
