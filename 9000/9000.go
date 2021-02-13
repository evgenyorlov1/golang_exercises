package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	s := strings.Split(a, "")
	sort.Strings(s)
	a = strings.Join(s, "")
	fmt.Println(a)
}