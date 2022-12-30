package main

import "fmt"

func main() {
	// Deklarasikan variabel nim dan platNomor
	var nim, platNomor int

	// Permintaan input nim dan plat nomor kepada user
	fmt.Print("Masukkan 1 angka belakang nim: ")
	fmt.Scan(&nim)
	fmt.Print("Masukkan plat nomor mobil: ")
	fmt.Scan(&platNomor)

	// Cek apakah plat nomor valid (tidak 0 dan tidak lebih dari 9999)
	if platNomor == 0 || platNomor > 9999 {
		fmt.Println("Plat nomor tidak valid.")
		return
	}

	// Cek apakah plat nomor sesuai dengan nim
	if (nim%2 == 0 && platNomor%2 == 0) || (nim%2 == 1 && platNomor%2 == 1) {
		// Jika sesuai, maka mobil diperbolehkan lewat
		fmt.Println("Mobil diperbolehkan lewat.")
	} else {
		// Jika tidak sesuai, maka mobil tidak diperbolehkan lewat
		fmt.Println("Mobil tidak diperbolehkan lewat.")
	}
}
