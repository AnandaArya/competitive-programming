package main

import "fmt"

func main() {
	var input, tahun, bulan, minggu, hari int

	fmt.Scan(&input)
	tahun = input / 360
	bulan = (input % 360) / 30
	minggu = ((input % 360) % 30) / 7
	hari = ((input % 360) % 30) % 7

	fmt.Print(tahun, bulan, minggu, hari)

}
