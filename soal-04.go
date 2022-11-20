package main

import "fmt"

func main() {
	var r, t, v float64
	const phi = 3.14
	fmt.Scanln(&r, &t)

	v = (phi * (r * r) * t) / 3
	fmt.Print(v)
}
