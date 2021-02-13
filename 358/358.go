package main

import "fmt"

func main() {
	var res int
	var n string
	fmt.Scan(&n)
	for _, char := range n {
		switch char {
		case '0': res += 6
		case '1': res += 2
		case '2': res += 5
		case '3': res += 5
		case '4': res += 4
		case '5': res += 5
		case '6': res += 6
		case '7': res += 3
		case '8': res += 7
		case '9': res += 6
		}	
	}
	fmt.Println(res)
}