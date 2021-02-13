package main

import (
	"fmt"
	"strings"
	"bufio"
  	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var str string
	str, _ = in.ReadString('\n')
	str = strings.ToLower(str)
	var c string
	fmt.Scan(&c)

	var count int
	for _, val := range str {
		if val == rune(c[0]) {
			count += 1
		}
	}
	fmt.Println(count)
}