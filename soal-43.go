package main

import "fmt"

func main() {
	var nilai1, nilai2, nilai3 float64
	var hasil bool

	for {
		fmt.Scan(&nilai1, &nilai2, &nilai3)
		if nilai1 < 0 || nilai2 < 0 || nilai3 < 0 {
			break
		}
		hasil = nilai1 > 50 && nilai2 > 50 && nilai3 > 50
		fmt.Println(hasil)
	}

}
