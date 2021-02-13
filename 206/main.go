package main

import "fmt"

func main()  {
	var k, w, a1, b1, a2, b2, a3, b3 int
	fmt.Scan(&k, &w, &a1, &b1, &a2, &b2, &a3, &b3)
	
	var flag bool = false

	if ((b1 >= k) && (a1 <= w)) {
		flag = true
	} 
	if ((b2 >= k) && (a2 <= w)) {
		flag = true
	}
	if ((b3 >= k) && (a3 <= w)) {
		flag = true
	}
	if ((b1 + b2 >= k) && (a1 + a2 <= w)) {
		flag = true
	}
	if ((b1 + b3 >= k) && (a1 + a3 <= w)) {
		flag = true
	}
	if ((b2 + b3 >= k) && (a2 + a3 <= w)) {
		flag = true
	}
	if ((b1 + b2 + b3 >= k) && (a1 + a2 + a3 <= w)) {
		flag = true
	}
	
	if flag {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}