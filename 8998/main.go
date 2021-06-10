package main

import "fmt"

func main()  {
	var ln string
	fmt.Scan(&ln)

	var yg rune = 999
	for _, c := range ln {
		if yg > c {
			yg = c
		}
	}

	count := 0
	for _, c := range ln {
		if c ==  yg {
			count += 1
		}
	}
	fmt.Printf("%c ", yg)
	fmt.Print(count)
}