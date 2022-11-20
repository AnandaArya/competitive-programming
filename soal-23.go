package main

import "fmt"

func main() {
	var a, b, jumlahAtas int
	fmt.Print("masukan bilangan ke 1 2 3 4 : ")
	fmt.Scan(&a)
	fmt.Scan(&b)

	jumlahAtas = a + 0 + b
	fmt.Print(jumlahAtas)
}
