package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type Vertex struct {
	X, Y float64
}

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	v := Vertex{3, 4}
	i = v
	describe(i)
}
