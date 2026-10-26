package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// bufio.NewReader wraps os.Stdin so it can be read a piece at a time, and
	// ReadString('\n') hands back everything up to and including the first
	// newline: one line of input.
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	parts := strings.Fields(strings.TrimSpace(line))
	nums := make([]int, 0, len(parts))
	for _, p := range parts {
		n, _ := strconv.Atoi(p)
		nums = append(nums, n)
	}

	winner := nums[0]
	for _, n := range nums {
		if winner < n {
			winner = n
		}
	}

	fmt.Println(winner)
}
