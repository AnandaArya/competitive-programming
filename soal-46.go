package main

import "fmt"

func main() {
	var sum int

	sum = 0

	for i := 1; i <= 99; i += 2 {
		sum += i
	}

	fmt.Print(sum)
}
