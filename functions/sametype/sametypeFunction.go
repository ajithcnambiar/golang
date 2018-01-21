package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Printf("add these vars %d \n", add(8, 9))
}
