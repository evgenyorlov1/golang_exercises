package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int64
	var k int
	fmt.Scan(&n, &k)

	num := big.NewInt(n)

	var factor int64 = 10
	for i := 1; i < k; i++ {
		temp := num
		num.Mul(num, big.NewInt(factor))
		num.Add(temp, num)
		fmt.Println(num.String())
		factor *= 10
	}

	total := new(big.Int).Mul(num, num)
	fmt.Print(total.String())
}
