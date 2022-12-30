package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanln(&a, &b)

	if b%a == 0 {
		fmt.Println("Ya,", a, "faktor", b)
	} else {
		fmt.Println(a, "bukan faktor dari", b)
	}
}
