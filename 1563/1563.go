package main

import (
	"fmt"
	"bufio"
  	"os"
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
	reader := bufio.NewReader(os.Stdin)
	
	var n, res int
	fmt.Scan(&n)
	
	shops := []string{}

	var shop string
	res = 0
	for i:=0; i<n; i++ {
		shop, _ = reader.ReadString('\n')		
		if stringInArray(shop, shops) {
			res += 1
		} else {
			shops = append(shops, shop)
		}
	}

	fmt.Println(res)
}