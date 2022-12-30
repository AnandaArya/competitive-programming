package main

import (
	"fmt"
)

func main() {
	var totalSiswa, kandidatA, kandidatB int
	fmt.Scanln(&totalSiswa, &kandidatA, &kandidatB)

	totalkandidat := kandidatA + kandidatB

	prosentasekandidat := float64(totalkandidat) / float64(totalSiswa)
	if prosentasekandidat >= 0.6 {
		if kandidatA > kandidatB {
			fmt.Println("Kandidat A menang")
		} else if kandidatB > kandidatA {
			fmt.Println("Kandidat B menang")
		}
	} else {
		fmt.Println("kandidatan ketua kelas tidak sah")
	}
}
