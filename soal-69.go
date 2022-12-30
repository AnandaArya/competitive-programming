package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	if (b == a+1 && c == b+1) || (b == a-1 && c == b-1) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
