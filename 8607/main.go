package main

import "fmt"

func main()  {
	var n string
	var res int
	fmt.Scan(&n)

	for _, v := range n {
		res += int(v - '0') * int(v - '0')
	}

	fmt.Print(res)
}