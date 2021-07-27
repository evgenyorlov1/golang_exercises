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

	c := new(big.Int).Add(a, b)
	fmt.Print(c.String())
}
