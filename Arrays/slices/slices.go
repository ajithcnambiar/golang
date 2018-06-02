package main

import "fmt"

func main() {
	primes := [6]int{3, 4, 5, 6, 7, 8}

	var s []int = primes[1:4]
	fmt.Println(s)
}
