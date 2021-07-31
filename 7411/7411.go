package main

import "fmt"

func main() {
	var a, b, c, d, e rune
	fmt.Scanf("%c%c%c%c%c", &a, &b, &c, &d, &e)

	if a == 'x' {
		switch b {
		case '+':
			fmt.Print(e - c)
		case '-':
			fmt.Print(e + c - '0' - '0')
		}
	}
	if c == 'x' {
		switch b {
		case '+':
			fmt.Print(e - a)
		case '-':
			fmt.Print(a - e)
		}
	}
	if e == 'x' {
		switch b {
		case '+':
			fmt.Print(a + c - '0' - '0')
		case '-':
			fmt.Print(a - c)
		}
	}
}
