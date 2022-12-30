package main

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	genap := 0
	ganjil := 0

	if a%2 == 0 {
		genap++
	} else {
		ganjil++
	}

	if b%2 == 0 {
		genap++
	} else {
		ganjil++
	}

	if c%2 == 0 {
		genap++
	} else {
		ganjil++
	}

	if d%2 == 0 {
		genap++
	} else {
		ganjil++
	}

	fmt.Print(genap, ganjil)
}
