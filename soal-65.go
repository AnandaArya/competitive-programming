package main

import (
	"fmt"
)

func main() {
	var n, input, i int

	fmt.Scan(&n)

	i = 1

	totalTabungan := 0

	for i <= n {
		fmt.Scan(&input)
		totalTabungan += input
	}

	fmt.Print(totalTabungan)

}
