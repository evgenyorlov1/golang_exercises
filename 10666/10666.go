package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	mp := make(map[string]int)

	for scanner.Scan() {
		inpt := scanner.Text()
		if inpt == "" {
			break
		}
		vals := strings.Split(inpt, " ")
		if _, ok := mp[vals[0]]; ok {
			v, _ := strconv.Atoi(vals[1])
			mp[vals[0]] += v
		} else {
			v, _ := strconv.Atoi(vals[1])
			mp[vals[0]] = v
		}
	}

	for k, v := range mp {
		fmt.Println(k, " ", v)
	}
}
