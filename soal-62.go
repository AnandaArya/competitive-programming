package main

import "fmt"

func main() {
	var sampo, hujan string
	fmt.Scan(&sampo, &hujan)

	if sampo == "ya" && hujan == "tidak" {
		fmt.Println("tidak pergi ke minimarket")
	} else if sampo == "tidak" || hujan == "tidak" {
		fmt.Println("tidak pergi ke minimarket")
	} else {
		fmt.Println("pergi ke minimarket")
	}
}
