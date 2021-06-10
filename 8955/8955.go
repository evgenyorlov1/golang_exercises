package main

import "fmt"

func main()  {
	var n, val int
	fmt.Scan(&n)

	count := 0

	var arr []int
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		if val > 0 {
			count++
			arr = append(arr, val)
		}
	}

	if count == 0 {
		fmt.Print("NO")
		return
	}

	fmt.Println(count)
	for _, v := range arr {
		fmt.Print(v, " ")
	}
}