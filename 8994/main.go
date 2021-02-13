package main

import (
	"fmt"
	"strings"
)

func main()  {
	var a, b string
	fmt.Scan(&a, &b)
	
	for _, v := range b {
		if strings.ContainsRune(a, v) == false {
			fmt.Print("No")		
			return
		}
	}
	fmt.Print("Ok")
}