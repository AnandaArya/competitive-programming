package main

import "fmt"

func main() {
	// Deklarasikan variabel iv
	var pokemon int

	// Permintaan input pokemon kepada user
	fmt.Print("Masukkan nilai pokemon pokemon: ")
	fmt.Scan(&pokemon)

	// Cek apakah pokemon sesuai dengan range yang ditentukan
	if pokemon >= 0 && pokemon <= 31 {
		if pokemon == 0 {
			fmt.Println("No Good")
		} else if pokemon >= 1 && pokemon <= 15 {
			fmt.Println("Decent")
		} else if pokemon >= 16 && pokemon <= 25 {
			fmt.Println("Pretty Good")
		} else if pokemon >= 26 && pokemon <= 29 {
			fmt.Println("Very Good")
		} else if pokemon == 30 {
			fmt.Println("Fantastic")
		} else {
			fmt.Println("Best")
		}
	} else {
		fmt.Println("Masukan harus lebih besar dari 0 dan kurang dari 31")
	}
}
