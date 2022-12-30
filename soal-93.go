package main

import "fmt"

func main() {
	// Meminta input angka tebakan pertama dari Shinichi
	var shinichi int
	fmt.Print("Masukkan angka tebakan Shinichi: ")
	fmt.Scan(&shinichi)

	// Meminta input angka tebakan kedua dari Heiji
	var heiji int
	fmt.Print("Masukkan angka tebakan Heiji: ")
	fmt.Scan(&heiji)

	// Menghitung hasil perselisihan dari kedua angka tebakan
	hasil := shinichi - heiji

	// Menentukan siapa yang menang sesuai dengan aturan permainan
	if hasil > 0 {
		if hasil%2 == 0 && hasil <= 45 {
			fmt.Println("Shinichi menang")
		} else {
			fmt.Println("Shinichi dan Heiji kalah")
		}
	} else if hasil < 0 {
		if hasil%2 != 0 && hasil >= -16 {
			fmt.Println("Heiji menang")
		} else {
			fmt.Println("Shinichi dan Heiji kalah")
		}
	} else {
		fmt.Println("Shinichi dan Heiji seri")
	}
}
