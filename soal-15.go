package main

import (
	"fmt"
	"math"
)

func main() {
	var b float64
	var a int
	var hasil bool

	fmt.Scan(&a, &b)

	b = math.Sqrt(b)

	hasil = a == int(b)

	fmt.Print(hasil)
}
