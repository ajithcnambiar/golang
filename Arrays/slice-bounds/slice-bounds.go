package main

import "fmt"

func main() {
	s := []int{4, 2, 9, 3, 7}

	s = s[1:5]
	fmt.Println(s)

	s1 := s[:2]
	fmt.Println(s1)

	s = s[0:]
	fmt.Println(s)
}
