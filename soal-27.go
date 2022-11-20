package main

import "fmt"

func main() {
	var b1, b2, b3, kg, g int

	fmt.Scan(&b1, &b2, &b3)

	kg = (b1 + b2 + b3) / 1000
	g = (b1 + b2 + b3) % 1000

	fmt.Print(kg, g)
}
