package main

import (
	"fmt"
)

// square should return n multiplied by itself.
func square(n int) int {
	return n * n
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(square(n))
}
