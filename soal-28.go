package main

import "fmt"

func main() {
	var adik, kakak int
	var hasil bool

	fmt.Scan(&adik, &kakak)

	hasil = adik == kakak || kakak-adik%3 == 0 || adik-kakak%3 == 0

	fmt.Print(hasil)
}
