package main

import "fmt"

func main() {
	var a, b, i int
	var hasil bool
	fmt.Print("Cek apakah nilai dari bilangan pertama dan adalah faktor dari bilangan ke 2 : ")

	fmt.Scan(&a, &b)

	fmt.Println("factor dari ", b, "adalah :")

	for i = 1; i < b; i++ {
		if b%i == 0 {

			fmt.Println(i)
			if i == a {
				hasil = true
			}
		}
	}

	fmt.Print("hasilnya adalah ", hasil)
}
