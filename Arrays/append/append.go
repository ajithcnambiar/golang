package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 13)
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("slice = %v, cap = %v, len = %v\n", s, cap(s), len(s))
}
