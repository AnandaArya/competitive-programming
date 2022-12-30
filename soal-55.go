package main

import "fmt"

func main() {

	var x, bronze, silver, gold int

	fmt.Scan(&x)
	gold = x / 9
	silver = x % 9 / 3
	bronze = x % 3
	// gold = ((x / 3) / 3) % 3
	// silver = (x / 3) % 3
	// bronze = x % 3
	fmt.Print(gold, silver, bronze)

}
