package main

import "fmt"

func main() {
	var adik, kakak, jumlah int
	var hasil bool

	fmt.Scan(&adik, &kakak)

	jumlah = kakak - adik

	hasil = jumlah%2 == 0

	fmt.Print(hasil)
}
