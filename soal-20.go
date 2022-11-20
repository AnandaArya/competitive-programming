package main

import "fmt"

func main() {
	var d1, d2, d3 int
	var h1, h2, h3, hasil bool

	fmt.Print("masukan inputan : ")
	fmt.Scan(&d1, &d2, &d3)

	h1 = d1%2 == 0
	h2 = d2%2 == 0
	h3 = d3%2 == 0

	hasil = h1 && h2 && h3
	fmt.Print(hasil)
}
