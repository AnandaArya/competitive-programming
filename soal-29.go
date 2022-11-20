package main

import "fmt"

func main() {
	var input, d1, d2, d3, d4 int
	var diskon, cashback, voucher bool

	fmt.Print("masukan inputan : ")
	fmt.Scan(&input)

	d1 = input / 1000
	d2 = (input / 100) % 10
	d4 = input % 10

	diskon = ((d2*10)+d3)%2 == 0
	cashback = (d1+d3)%d4 == 0
	voucher = diskon != cashback

	fmt.Print(diskon, cashback, voucher)
}
