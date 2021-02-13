package main

import "fmt"
import "os"
import "bufio"

func main() {
	var word string
	in := bufio.NewReader(os.Stdin)
	word, _ = in.ReadString('\n')

	fmt.Print(word, "\n", len(word) - 1)
}