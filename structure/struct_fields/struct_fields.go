package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{3, 4}
	fmt.Printf("X is %v \n", v.X)
	fmt.Printf("Y is %v ", v.Y)
}
