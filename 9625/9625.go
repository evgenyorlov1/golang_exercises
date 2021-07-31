package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str = scanner.Text()
	fmt.Println(strings.ToUpper(str))
}
