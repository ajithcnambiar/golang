package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	p := &Vertex{3, 4}
	fmt.Printf("Before Scaling Abs = %v \n", p.Abs())

	p.Scale(10)
	fmt.Printf("After Scaling Abs = %v \n", p.Abs())
}
