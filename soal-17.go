package main

import "fmt"

func main() {
	var a, b int
	var hasil bool

	fmt.Scan(&a, &b)

	hasil = b%a == 0 || b == a

	fmt.Print(hasil)

}
