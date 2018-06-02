package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a := make([]int, 0)
	i := -1
	return func() int {
		i++
		if i == 0 {
			a = append(a, 0)
		} else if i == 1 {
			a = append(a, 1)
		} else {
			a = append(a, a[i-2]+a[i-1])
		}
		return a[i]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
