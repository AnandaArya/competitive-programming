package main

import "fmt"

func main() {

	var suwir, cakue, ati, telur bool
	var total int

	fmt.Scan(&suwir, &cakue, &ati, &telur)

	total = 6000

	if suwir {
		total += 3000
	}

	if cakue {
		total += 1500
	}

	if ati {
		total += 4500
	}

	if telur {
		total += 4000
	}

	fmt.Print(total)

}
