package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 3}
	v3 = Vertex{Y: 4}
	v4 = &Vertex{5, 6}
)

func main() {
	fmt.Printf("v1 = %v v2 = %v v3 = %v v4 = %v", v1, v2, v3, v4)
}
