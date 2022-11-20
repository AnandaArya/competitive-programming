package main

import "fmt"

func main() {
	var a int
	var hasil bool

	fmt.Scan(&a)

	hasil = a%2 == 1

	fmt.Print(hasil)
}
