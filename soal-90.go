package main

import "fmt"

func main() {
	var heiji int
	fmt.Print("Masukkan angka tebakan pertama Heiji: ")
	fmt.Scan(&heiji)

	var shinichi int
	fmt.Print("Masukkan angka tebakan kedua Shinichi: ")
	fmt.Scan(&shinichi)

	if heiji > shinichi {
		fmt.Printf("Hasil perselisihan: %d\n", heiji-shinichi)
	} else {
		fmt.Printf("Hasil perselisihan: %d\n", shinichi-heiji)
	}
}
