package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	sum := 0
	for k := 1; k <= n; k++ {
		sum += k
	}
	fmt.Println(sum)
}
