package main

import (
	"fmt"
)

func naked(a, b int) (x, y int) {
	x = a + b
	y = a - b
	return
}

func main() {
	x, y := naked(55, 16)
	fmt.Printf("naked function returns: %d %d", x, y)
}
