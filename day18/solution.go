package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"math"
)

var world map[int]bool = make(map[int]bool, 0)

var coordinates [][]int = make([][]int, 0)

func exists(x, y, z int) bool {
	_, ok := world[x + y*1000 + z*1000000]
	return ok
}

func main() {
	b, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		b := strings.Split(line, ",")
		l := 0
		xyz := make([]int, 0)
		for i := 0; i<3; i++ {
			v, _ := strconv.Atoi(b[i])
			xyz = append(xyz, v)
			l += v * int(math.Pow(10, float64(i*3)))
		}
		coordinates = append(coordinates, xyz)
		world[l] = true
	}

	score := 0
	for _, c := range coordinates {
		score += 6

		if exists(c[0]-1, c[1], c[2]) {
			score--
		}
		if exists(c[0]+1, c[1], c[2]) {
			score--
		}
		if exists(c[0], c[1]-1, c[2]) {
			score--
		}
		if exists(c[0], c[1]+1, c[2]) {
			score--
		}
		if exists(c[0], c[1], c[2]-1) {
			score--
		}
		if exists(c[0], c[1], c[2]+1) {
			score--
		}
	}
	fmt.Println(score)
}
