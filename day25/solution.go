package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var digits = [...]string{"2", "1", "0", "-", "="}
var sizes = [...]int{2, 1, 0, -1, -2}

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

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func toDec(x string) int {
	result := 0
	for index, c := range x {
		v := 0
		switch c {
		case '=':
			v = -2
		case '-':
			v = -1
		case '1':
			v = 1
		case '2':
			v = 2
		}
		result += v * pow(5, len(x)-1-index)
	}
	return result
}

func upperBound(n int) int {
	if n < 0 {
		return 0
	}
	return 2*pow(5, n) + upperBound(n-1)
}

func lowerBound(n int) int {
	if n < 0 {
		return 0
	}
	return -2*pow(5, n) + lowerBound(n-1)
}

func toSnafu(dec int) string {
	n := 0
	for !(dec >= lowerBound(n) && dec <= upperBound(n)) {
		n++
	}

	result := ""
	for n >= 0 {
		for index, size := range sizes {
			p := size * pow(5, n)
			if dec-p <= upperBound(n-1) && dec-p >= lowerBound(n-1) {
				dec -= p
				result += digits[index]
				break
			}
		}
		n--
	}

	return result
}

func main() {
	b, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(b), "\n")
	var total = 0
	for _, line := range lines {
		total += toDec(line)
	}
	fmt.Println(total)

	fmt.Println(toSnafu(total))
	// fmt.Println(upperBound(0), lowerBound(0))
}
