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
	fmt.Print(new(big.Int).Exp(a, b, nil))
}
