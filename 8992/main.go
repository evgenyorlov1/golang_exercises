package main

import (
	"fmt"
	"bufio"
  	"os"
)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
	line := scanner.Text()

	fmt.Printf("%c", line[0])
	for i := 1; i < len(line); i++ {
		if line[i - 1] == line[i] {
			continue
		}
		fmt.Printf("%c", line[i])
	}

}