package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	fmt.Printf("sum of 3 and 4 is %d \n", add(3, 4))
	fmt.Printf("sqrt of 9 is %g \n", sqrt(9))
}
