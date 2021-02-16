package main

import (
	"fmt"
	"os"
	"bufio"
	"unicode"
)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
	line := scanner.Text()


	m := make(map[rune]int)
	for _, v := range line {
		if unicode.IsLetter(v) {
			m[v] += 1
			fmt.Printf("%c ", v)
		}
	}
	// var max int = -99
	for k, _ := range m {
		fmt.Printf("%c ", k)
		// fmt.Print(v, " ")
		// if v > max {
			// max = v
		// }
	}
	
	// result := []rune{}
	// for k, v := range m {
	// 	if v == max {
	// 		result = append(result, k)
	// 	}
	// }
	
	// fmt.Println(max)
	// for _, v := range result {
	// 	fmt.Printf("%c", v)
	// 	fmt.Print(" ")
	// }
}