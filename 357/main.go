package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func numbers(s string) []float64 {
    var n []float64
    for _, f := range strings.Fields(s) {
        i, err := strconv.ParseFloat(f, 2)
        if err == nil {
            n = append(n, i)
        }
    }
    return n
}

func main()  {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	
	var max float64 = -100.00
	for _, v := range numbers(scanner.Text()) {
		if v > max {
			max = v
		}
	}
	fmt.Printf("%.2f", max)
}