package main

import "fmt"

func main() {
	var total, hasil int
	fmt.Scan(&total)

	if total >= 275000 {
		hasil = total - (total * 5 / 100)
		fmt.Println("Total yang harus dibayar: ", hasil)
	} else {
		hasil = total
		fmt.Println("Total yang harus dibayar: ", hasil)
	}

}
