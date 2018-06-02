package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("cap = %v len = %v slice = %v", cap(s), len(s), s)

	if s == nil {
		fmt.Println("slice is nil")
	}
}
