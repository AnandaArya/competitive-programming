package main

import "fmt"

func main() {
	var x, y int
	var z float64
	fmt.Scan(&x, &y)
	z = ((3 * float64(x)) * (6 * float64(y))) / (4 * float64(y))
	fmt.Print(z)

}

// program kuadrat
// { program menghitung kuadrat }

// kamus
// x, i, jumlah: integer
// algoritma

// 	input(x)
// 	jumlah <- 0
// 	for i <- 1, to x do
// 		jumlah = jumlah + i * i
// 	endfor

// 	output(jumlah)
// endprogram
