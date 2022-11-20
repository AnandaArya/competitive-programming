package main

import "fmt"

func main() {

	// inputan bilangan genap akan berhenti apabila inputan ganjil
	var input int
	var sum int

	sum = 0

	for input%2 == 0 {
		fmt.Print("masukan nilai input :")
		fmt.Scan(&input)
		sum = sum + input
	}

	fmt.Print(sum)

}
