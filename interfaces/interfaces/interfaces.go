package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

// Types must implement the interface method(s)
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-1 * f)
	} else {
		return float64(f)
	}

}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)

	v := Vertex{3, 4}

	fmt.Printf("MyFloat  %v\n", f.Abs())

	fmt.Printf("Vertex  %v\n", v.Abs())

	a = &f
	fmt.Printf("MyFloat retrieved via interface %v\n", a.Abs())

	a = &v
	fmt.Printf("Vertex retrieved via interface %v\n", a.Abs())

}
