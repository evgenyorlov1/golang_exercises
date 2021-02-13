package main

import "bufio"
import "os"
import "fmt"

func main() {
	var res int
	in := bufio.NewReader(os.Stdin)
	n, _ := in.ReadString('\n')

	for _, char := range n {
		switch char {
		case 'A': res += 1
		case 'E': res += 1
		case 'I': res += 1
		case 'O': res += 1
		case 'U': res += 1
		case 'Y': res += 1
		}	
	}
	fmt.Println(res)
}