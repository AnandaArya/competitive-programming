package main

import "fmt"

func main() {
	var input string

	var hasil bool
	fmt.Scan(&input)

	if input == "a" || input == "i" || input == "u" || input == "e" || input == "o" {
		hasil = false
	}

	fmt.Print(hasil)
}

// hasil = input == "a" || input == "i" || input == "u" || input == "e" || input == "o"
