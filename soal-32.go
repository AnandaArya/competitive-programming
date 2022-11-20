package main

import (
	"fmt"
	"strings"
)

func main() {
	var input, hasil string

	fmt.Scan(&input)

	hasil = strings.ToLower(input)
	fmt.Print(hasil)
}
