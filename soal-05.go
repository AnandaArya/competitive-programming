package main

import "fmt"

func main() {
	var r, L float64
	const phi = 3.14
	fmt.Scanln(&r)

	L = 4 * phi * (r * r)
	fmt.Print(L)
}
