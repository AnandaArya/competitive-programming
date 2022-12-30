package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)
	if (a+b+c == c+d+e) && ((b+c+d)%(a+e) == 0) {
		fmt.Println("cashback")
		fmt.Println("diskon")
	} else if (b+c+d)%(a+e) == 0 {
		fmt.Println("diskon")
	} else if a+b+c == c+d+e {
		fmt.Println("cashback")
	}
}
