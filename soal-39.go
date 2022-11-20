package main

import (
	"fmt"
)

func main() {
	var input, sum, volume int
	var result bool

	fmt.Scan(&volume)

	sum = 0

	for sum < volume {
		fmt.Scan(&input)
		sum += input
		if sum < volume {
			result = false
			fmt.Println(result)
		} else {
			result = true
			fmt.Println(result)
		}
	}
}
