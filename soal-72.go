package main

import (
	"fmt"
	"math"
)

func main() {
	var busCapacity, penumpang float64

	fmt.Scanln(&busCapacity, &penumpang)

	var totalBus = penumpang / busCapacity

	totalBus = math.Ceil(totalBus)

	fmt.Print(totalBus)
}
