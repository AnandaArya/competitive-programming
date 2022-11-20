package main

import "fmt"

func main() {

	var n, m int
	var berhenti bool
	fmt.Scan(&n, &m)

	berhenti = false
	for !berhenti {
		n = n + 1
		fmt.Println(n)
		berhenti = n == m
	}

}

// var n, m int

// 	fmt.Scan(&n, &m)

// 	i := n
// 	for next := true; next; next = i <= m {
// 		fmt.Println(i)
// 		i++
// 	}
