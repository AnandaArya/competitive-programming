package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)

	var persentase = (b / a) * 100

	if math.Round(persentase) >= 50 && math.Round(persentase) <= 75 {
		fmt.Println("berangkat")
	} else {
		fmt.Println("tidak berangkat")
	}
}
