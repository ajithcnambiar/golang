package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)
}

func printSlice(a []int) {
	fmt.Printf("len = %v, cap = %v, slice = %v\n", len(a), cap(a), a)
}
