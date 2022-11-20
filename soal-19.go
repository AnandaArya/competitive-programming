package main

import "fmt"

func main() {
	var input, a, b, c, d int

	fmt.Print("masukan inputan : ")
	fmt.Scan(&input)

	a = input / 1000
	b = input / 100 % 10
	c = input % 100
	d = input % 10

	input = a + b + c + d

	fmt.Print(input)

}
