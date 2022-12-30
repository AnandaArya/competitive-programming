// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var input string

// 	fmt.Scan(&input)

// 	if len(input) >= 20 {
// 		pesanRahasia := string(input[0]) + string(input[4]) + string(input[8]) + string(input[12])

// 		fmt.Println(pesanRahasia)
// 	}
// }

package main

import "fmt"

func main() {
	var string1 string

	fmt.Scan(&string1)

	if len(string1) < 20 {

	} else {
		fmt.Printf("%c", string1[0])
		fmt.Printf("%c", string1[4])
		fmt.Printf("%c", string1[8])
		fmt.Printf("%c", string1[12])
	}
}
