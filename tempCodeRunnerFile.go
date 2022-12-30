package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	var persentase = (b / a) * 100

	if persentase >= 50 && persentase <= 75 {
		fmt.Print("berangkat")
	} else {
		fmt.Print("tidak berangkat")
	}
}
