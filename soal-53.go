package main

import "fmt"

func main() {

	var input int
	var hasil bool

	fmt.Scan(&input)
	hasil = true && input != 1

	for i := 2; i < input; i++ {
		hasil = hasil && input%i != 0
	}

	fmt.Print(hasil)

}
