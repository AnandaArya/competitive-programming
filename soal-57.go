package main

import "fmt"

func main() {

	var a, b, c int

	fmt.Scan(&a, &b, &c)

	if a < c {
		a, c = c, a
	}

	if a < b {
		a, b = b, a
	}

	if b < c {
		b, c = c, b
	}

	fmt.Print(a, b, c)

}
