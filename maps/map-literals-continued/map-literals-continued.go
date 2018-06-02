package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell labs": {40.68, -74.399},
	"Google":    {37.422, -122.08},
}

func main() {
	fmt.Println(m)
}
