package main

import "fmt"

func main() {
	var a, b rune

	fmt.Scan(&a, &b)

	var hasil = false

	if (a == 'B' && b == 'A') || (b == 'B' && a == 'A') || (a == 'B' && b == 'D') || (b == 'B' && a == 'D') || (a == 'B' && b == 'C') || (b == 'B' && a == 'C') || (a == 'C' && b == 'E') || (b == 'C' && a == 'E') {
		hasil = true
	}

	fmt.Print(hasil)
}
