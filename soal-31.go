package main

import "fmt"

func main() {
	var input1, input2, d1, d2, hasil int

	fmt.Scan(&input1, &input2)
	d1 = input1 / 100
	d2 = input1 % 100

	hasil = d1*1000 + input2*100 + d2

	fmt.Print(hasil)
}
