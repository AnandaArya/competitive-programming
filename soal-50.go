package main

import "fmt"

func main() {
	var input, i, sum int

	sum = 0
	for {
		fmt.Scan(&input)
		if input%2 == 1 {
			break
		}

		sum += input
		i++
	}
	fmt.Print(sum, i)
}
