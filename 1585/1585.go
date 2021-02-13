package main

import "fmt"

func main() {
	var a, e, l, m int
	fmt.Scan(&a)
	
	var pinguin string
	
	for i:=0; i<a; i++ {
		fmt.Scan(&pinguin)
		if pinguin[0] == 'E' {
			e += 1
		}
		if pinguin[0] == 'L' {
			l += 1
		}
		if pinguin[0] == 'M' {
			m += 1
		}
		fmt.Scan(&pinguin)
	}
	
	if e >= l && e >= m {
		fmt.Println("Emperor Penguin")	
		return
	}
	if l >= e && l >= m {
		fmt.Println("Little Penguin")	
		return
	}
	if m >= l && m >= e {
		fmt.Println("Macaroni Penguin")	
		return
	}
}