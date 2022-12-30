package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)

	var mid int
	if a < b && a < c {
		if b < c {
			mid = b
		} else {
			mid = c
		}
	} else if b < a && b < c {
		if a < c {
			mid = a
		} else {
			mid = c
		}
	} else {
		if a < b {
			mid = a
		} else {
			mid = b
		}
	}

	fmt.Println(mid)
}
