package main

import "golang.org/x/tour/pic"

// Pic Sample export
func Pic(dx, dy int) [][]uint8 {
	outer := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		outer[i] = make([]uint8, dx)
	}
	return outer
}

func main() {
	pic.Show(Pic)
}
