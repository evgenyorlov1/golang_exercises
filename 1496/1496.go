package main

import (
	"fmt"
	// "bufio"
  	// "os"
)

func stringInArray(a string, list []string) bool {
	for _, b := range list {
		if a == b {
			return true
		}
	}
	return false
}

func main() {
	// reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Scanf("%d", &n)

	names := []string{}

	var name string
	for i:=0; i<n; i++ {
		// name, _ = reader.ReadString('\n')
		fmt.Scanln("%s", &name)
		if stringInArray(name, names) {
			fmt.Println(name)
		} else {
			names = append(names, name)
		}
	}
}