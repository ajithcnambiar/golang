package main

import (
	"fmt"
	"math/cmplx"
)

var (
	toBe           = false
	maxInt  uint64 = 1<<64 - 1
	z              = cmplx.Sqrt(-5 + 12i)
	unicode        = '\u0D05'
)

func main() {
	fmt.Printf("Type = %T and value = %v \n", toBe, toBe)
	fmt.Printf("Type = %T and value = %v \n", maxInt, maxInt)
	fmt.Printf("Type = %T and Value = %v \n", z, z)
	fmt.Printf("Type = %T and Value = %v", unicode, unicode)
}
