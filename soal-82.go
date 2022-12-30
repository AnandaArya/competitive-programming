package main

import "fmt"

func main() {

	var a, b string

	var hasil bool

	fmt.Scan(&a, &b)

	if (a == "B" && b == "A") || (b == "B" && a == "A") || (a == "B" && b == "C") || (b == "B" && a == "C") || (a == "A" && b == "D") || (b == "A" && a == "D") {
		hasil = true
	} else if (a == "C" && b == "E") || (b == "C" && a == "E") {
		hasil = true
	}

	fmt.Print(hasil)

}
