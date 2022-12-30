package main

import "fmt"

func main() {

	var i int
	var hasil bool

	for i = 1; i <= 10; i++ {
		hasil = i == 2 || i == 3 || i == 5 || i == 7
		fmt.Println(i, hasil)
	}
}
