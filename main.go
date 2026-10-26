package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	line = strings.TrimRight(line, "\r\n")

	val, err := strconv.Atoi(line)
	result := "ok"
	if err != nil {
		result = "bad"
		fmt.Println(result)
		return
	}

	fmt.Println(result, val)
}
