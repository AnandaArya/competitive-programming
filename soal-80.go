package main

import "fmt"

func main() {

	var a, b, c string

	fmt.Scan(&a, &b, &c)

	if (a == "yes" && b == "yes") || (a == "yes" && c == "yes") || (b == "yes" && c == "yes") {
		fmt.Print("lolos")
	} else {
		fmt.Print("tidak lolos")
	}

}
