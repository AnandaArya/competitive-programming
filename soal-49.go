package main

import "fmt"

func main() {
	var input string

	for {
		fmt.Scan(&input)
		if input == "ada" {
			fmt.Println("pencarian selesai")
			break
		}

		fmt.Println("cari beasiswa")
	}
}
