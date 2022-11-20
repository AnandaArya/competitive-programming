package main

import "fmt"

func main() {
	var a, b int

	var hasil bool = true
	fmt.Scan(&a, &b)

	if a < b && b%a == 0 {
		hasil = false
	}

	fmt.Print(hasil)
}
