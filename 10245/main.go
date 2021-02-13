package main

import "fmt"

func main()  {
	var n, m, s, e int
	fmt.Scan(&n, &m)

	names := []string{}
	var name string
	for i := 0; i < n; i++ {
		fmt.Scan(&name)
		names = append(names, name)
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&s, &e)	
		fmt.Printf("%c", names[s - 1][e])
	}
}