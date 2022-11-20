package main

import (
	"fmt"
)

func main() {
	var input1, input2 string

	fmt.Print("masukan inputan : ")
	fmt.Scan(&input1, &input2)

	input1 = input1[0:2] + input2 + input1[2:4]

	fmt.Print(input1)

}
