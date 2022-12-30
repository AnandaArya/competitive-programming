package main

import "fmt"

func main() {

	var a, b, c, d, total int

	fmt.Scan(&a, &b, &c, &d)

	total = a + b + c

	if d == total {
		fmt.Print("benar")
	} else {
		fmt.Print("salah")
	}

}
