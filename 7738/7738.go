package main

import "fmt"

func main() {
	var m, solution, total_time int
	var t, r string

	res := make(map[string]int)

	for {
		fmt.Scan(&m, &t)
		fmt.Println(m, " ", t)
		if m != -1 {
			fmt.Println("m != -1")
			if _, ok := res[t]; ok {
				res[t] += m + 20
			} else {
				res[t] += m
			}
			if r == "right" {
				solution += 1
				total_time += res[t]
			}
		} else {
			break
		}
	}

	fmt.Print(solution, " ", total_time)

}
