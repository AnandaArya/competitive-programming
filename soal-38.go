package main

import (
	"fmt"
)

func main() {
	var n, m, sum float64

	fmt.Scan(&n, &m)

	sum = 0.0

	for i := int(n); i <= int(m); i++ {
		sum += (4 / float64(i))
	}

	fmt.Printf("%.2f", sum)
}
