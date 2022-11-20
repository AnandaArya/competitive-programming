package main

import (
	"fmt"
	"strings"
)

func main() {
	var input1, input2 string
	var hasil bool

	fmt.Scan(&input1, &input2)
	hasil = strings.ToUpper(input1) == strings.ToUpper(input2) || strings.ToUpper(input1) == strings.ToLower(input2) || strings.ToLower(input1) == strings.ToUpper(input2)
	fmt.Print(hasil)
}
