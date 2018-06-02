package main

import "fmt"

func main() {
	names := [4]string{
		"Hello",
		"World",
		"is",
		"first",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]

	a[0] = "changed_a"
	b[1] = "changed_b"

	fmt.Println(a, b)
	fmt.Println(names)
}
