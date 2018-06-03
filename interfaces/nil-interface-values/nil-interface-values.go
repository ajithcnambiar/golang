package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(t I) {
	fmt.Printf("%v %T\n", t, t)
}
