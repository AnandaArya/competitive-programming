package main

import "fmt"

func main() {

	var a, b, c bool

	fmt.Scan(&a, &b, &c)

	if a && b && c {
		fmt.Print("berkemah")
	} else {
		fmt.Print("tidak berkemah")
	}

}
