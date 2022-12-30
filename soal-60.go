package main

import "fmt"

func main() {
	var a, b, c, d, e int

	fmt.Scan(&a, &b, &c, &d, &e)

	if a%2 == 1 && b%2 == 1 && c%2 == 1 && d%2 == 1 && e%2 == 1 {
		fmt.Print("ganjil")
	} else if a%2 == 0 && b%2 == 0 && c%2 == 0 && d%2 == 0 && e%2 == 0 {
		fmt.Print("genap")
	} else {
		fmt.Print("bukan ganjil atau genap")
	}
}
