package main

import (
	"fmt"
)

func main() {
	fmt.Println(factorial(3))
}

func factorial(x int) int {
	if x <= 1 {
		return 1
	}

	return x * factorial(x-1)
}
