package main

import "fmt"

func main() {
	var hp, attack, defense float64
	var iv float64
	fmt.Scan(&hp, &attack, &defense)

	iv = ((hp + attack + defense) / 45) * 100

	if hp < 0 || hp > 15 || attack < 0 || attack > 15 || defense < 0 || defense > 15 {
		fmt.Println("Not a pokemon")
		return
	}

	if iv < 0 || iv > 100 {
		fmt.Println("Not a pokemon")
		return
	}

	if iv < 48.90 {
		fmt.Println("Overall your pokemon has room for improvement as far as battling goes and 0 stars")
	} else if iv < 64.45 {
		fmt.Println("Overall your pokemon is pretty decent and 1 star")
	} else if iv < 80 {
		fmt.Println("Overall your pokemon is really strong and 2 stars")
	} else if iv < 97.8 {
		fmt.Println("Overall your pokemon looks like it can really battle with the best of them and 3 stars")
	} else if iv == 100 {
		fmt.Println("Overall your pokemon looks like it can really battle with the best of them and 4 stars")
	}
}
