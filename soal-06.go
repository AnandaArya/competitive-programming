package main

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c)
	d = (a + 1) * ((2 * b) + 2) * ((3 * c) + 3)

	fmt.Print(d)
}
