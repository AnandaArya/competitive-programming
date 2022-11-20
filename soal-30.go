package main

import "fmt"

func main() {
	var input, d1, d4 int
	var hasil bool

	fmt.Scan(&input)

	d1 = input / 1000

	d4 = input % 10

	hasil = d1 == d4

	fmt.Print(hasil)
}
