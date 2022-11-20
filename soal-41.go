package main

import "fmt"

func main() {
	var n, m, sum int

	fmt.Scan(&n, &m)

	sum = 0

	for i := n; i <= m; i++ {
		sum += i
	}

}
