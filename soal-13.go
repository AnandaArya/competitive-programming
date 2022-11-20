package main

import "fmt"

func main() {
	var input int
	var hasil bool

	fmt.Scan(&input)

	hasil = input < 0

	fmt.Print(hasil)

}
