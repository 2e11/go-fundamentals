package main

import (
	"errors"
	"fmt"
)

func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	if q, err := safeDivide(a, b); err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Printf("result: %d\n", q)
	}
}
