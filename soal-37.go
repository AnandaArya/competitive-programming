package main

import "fmt"

func main() {
	var x, y, n, total, sum, result int

	fmt.Scan(&x, &y, &n)

	total = x + y

	sum = 0

	for i := 1; i <= n; i++ {
		if i%total == 0 {
			sum += 1
		}
	}

	result = sum * 2

	fmt.Print(result)
}
