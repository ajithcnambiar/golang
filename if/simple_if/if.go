package main

import (
	"fmt"
)

func isGreater(x int) bool {
	if x > 5 {
		return true
	}
	return false
}

func main() {
	fmt.Println("4 is greater than 5: ", isGreater(4))
	fmt.Println("7 is greater than 5: ", isGreater(7))
}
