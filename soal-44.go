package main

import "fmt"

func main() {

	// 4000 + 6500 + 9000 + 11500
	var tabungan, minggu, sum, tambahan int

	fmt.Scan(&tabungan, &minggu)

	sum = 0
	tambahan = 2500
	for i := 1; i <= minggu; i++ {
		sum += tabungan + (tambahan * (i - 1))
	}

	fmt.Print(sum)
}
