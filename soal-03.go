package main

import "fmt"

// func main() {
// 	var x, d1 float64
// 	var xInteger int

// 	fmt.Print("Masukan nilai x : ")
// 	fmt.Scanln(&x)
// 	if x > 0 && x < 1000 {
// 		d1 = x / 100
// 		xInteger = int(d1)

// 		fmt.Print("keluaran nilai x adalah : ", d1)
// 		fmt.Print("keluaran nilai x adalah : ", d1, " setelah di convert jadi integer : ", xInteger)
// 	} else {
// 		fmt.Print("nilai x harus lebih besar dari 0 dan harus kurang dari 1000")
// 	}

// }

func main() {
	var x, d1, d2, d3 int

	fmt.Print("Masukan nilai x : ")
	fmt.Scanln(&x)
	if x > 0 && x < 1000 {
		d1 = x / 100
		d2 = x / 10 % 10
		d3 = x % 10
		fmt.Print("keluaran nilai x adalah : ", d1, d2, d3)
	} else {
		fmt.Print("nilai x harus lebih besar dari 0 dan harus kurang dari 1000")
	}

}
