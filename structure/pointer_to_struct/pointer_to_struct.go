package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{3, 4}
	p := &v
	fmt.Printf("Print by (*p).X %v \n", (*p).X)
	fmt.Printf("Print by p.X %v \n", p.X)

	p.X = 1e9
	fmt.Println(v)
}
