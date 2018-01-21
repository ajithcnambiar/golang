package main

import (
	"fmt"
)

func main() {
	x := 1
	sum := 0
	for x < 10 {
		sum += x
		x++
	}
	fmt.Println(sum)
}
