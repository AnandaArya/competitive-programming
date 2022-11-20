package main

import "fmt"

func main() {
	var a, b int
	var hasil bool = false

	fmt.Scan(&a, &b)

	if a < b && b%a == 0 {
		hasil = true
	}

	fmt.Print(hasil)
}
