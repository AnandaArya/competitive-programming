package main

import "fmt"

func main() {

	var a, b, total int

	fmt.Scan(&a, &b)

	if a%2 == 0 && b%2 == 0 {
		total = a * b
	} else if a%2 == 1 && b%2 == 1 {
		total = a + b
	} else {
		total = 0
	}

	fmt.Print(total)

}
