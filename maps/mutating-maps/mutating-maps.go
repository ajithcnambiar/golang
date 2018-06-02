package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["job"] = 1
	fmt.Println("The value: ", m["job"])

	m["job"] = 3
	fmt.Println("The value: ", m["job"])

	delete(m, "job")
	fmt.Println("The value: ", m["job"])

	v, ok := m["job"]

	fmt.Println("The value:", v, "Present?", ok)

	m["culture"] = 4
	v, ok = m["culture"]

	fmt.Println("The value:", v, "Present?", ok)
}
