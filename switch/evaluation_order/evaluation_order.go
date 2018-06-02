package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Print(" today")
	case today + 1:
		fmt.Print(" tomorrow")
	case today + 2:
		fmt.Print(" day after tomorrow")
	default:
		fmt.Println("Too far away")
	}
}
