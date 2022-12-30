package main

import "fmt"

func main() {
	var sehat, bekal, cuacaBagus bool

	fmt.Scan(&sehat, &bekal, &cuacaBagus)

	if sehat && bekal && cuacaBagus {
		fmt.Print("berkemah")
	} else {
		fmt.Print("tidak berkemah")
	}
}
