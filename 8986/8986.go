package main

import "fmt"
import "os"
import "bufio"

func main() {
	var phrase string
	var n, m int
	in := bufio.NewReader(os.Stdin)
	phrase, _ = in.ReadString('\n')
	fmt.Scan(&n, &m)
	fmt.Print(phrase[:n] + phrase[m+1:])
}