package main

import "fmt"

func main() {
	q := []int{3, 92, 2, 9, 8, 4}
	fmt.Println(q)

	r := []bool{true, false, true, false, true}

	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, false},
		{4, true},
		{5, true},
		{8, false},
	}

	fmt.Println(s)
}
