package main

import "fmt"

func main() {
	var n, m, x, y, i, totalBahan, totalInput, hasil int
	var berhenti bool

	fmt.Scan(&n, &m, &x, &y)

	totalBahan = n + m
	totalInput = x + y
	hasil = 0

	berhenti = false
	for !berhenti {
		i += 1
		if i%totalInput == 0 {
			hasil += 1
		}
		berhenti = i == totalBahan
	}

	fmt.Print(hasil)
}

// for i = 1; i <= totalBahan; i++ {
// 	if i%totalInput == 0 {
// 		hasil += 1
// 	}
// }
