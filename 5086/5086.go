package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)
	fmt.Scan(a)
	fmt.Scan(b)
	cmp := a.Cmp(b)
	switch cmp {
	case -1:
		fmt.Print("<")
	case 0:
		fmt.Print("=")
	case 1:
		fmt.Print(">")
	}
}
