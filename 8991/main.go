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

	for _, v := range line {
		if unicode.IsLetter(v) {
			fmt.Printf("%c", v)
			fmt.Printf("%c", v)
			continue
		}
		fmt.Printf("%c", v)
	}	
}