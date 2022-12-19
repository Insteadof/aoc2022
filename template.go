package main

import (
	"fmt"
	"os"
)

func strToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	b, _ := os.ReadFile("input.txt")

	line := string(b)

}
