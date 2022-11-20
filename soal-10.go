package main

import "fmt"

func main() {
	var input string
	var hasil bool

	fmt.Print("Masukan inputan 4 bilangan :")
	fmt.Scan(&input)

	hasil = input[0:1] == input[len(input)-1:]

	fmt.Print(hasil)
}
