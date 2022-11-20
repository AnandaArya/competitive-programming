package main

import "fmt"

func main() {
	var n, m, sum int

	fmt.Scan(&n, &m)
	for i := n; i <= m; i++ {
		sum += i
	}

	fmt.Print(sum)
}
