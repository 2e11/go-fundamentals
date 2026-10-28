package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	words := strings.Fields(strings.TrimSpace(line))
	counts := make(map[string]int)

	for _, val := range words {
		counts[val]++
	}
	different := len(counts)

	once := 0
	for _, val := range counts {
		if val == 1 {
			once++
		}
	}

	fmt.Println(different)
	fmt.Println(once)
}
