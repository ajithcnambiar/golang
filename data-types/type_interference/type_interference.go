package main

import (
	"fmt"
)

func main() {
	var i = 4
	j := i
	g := 3 + 56i
	fmt.Printf("Value  of interfered var = %d and type = %T", j, j)
	fmt.Printf("Value  of interfered var = %g and type = %T", g, g)
}
