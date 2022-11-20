package main

import "fmt"

func main() {
	var number int
	var word string

	fmt.Scan(&number, &word)

	for i := 1; i <= number; i++ {
		fmt.Print(word)

	}

}
