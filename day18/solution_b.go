package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"math"
)

var world map[int]int = make(map[int]int, 0)

var coordinates [][]int = make([][]int, 0)

var worldMin []int = []int { 1000, 1000, 1000 }
var worldMax []int = make([]int, 3)

func exists(x, y, z int) bool {
	_, ok := world[x + y*1000 + z*1000000]
	return ok
}

func min(arr1, arr2 []int) []int {
	result := make([]int, len(arr1))
	for i := 0; i<len(arr1); i++ {
		if arr1[i] < arr2[i] {
			result[i] = arr1[i]
		} else {
			result[i] = arr2[i]
		}
	}
	return result
}
func max(arr1, arr2 []int) []int {
	result := make([]int, len(arr1))
	for i := 0; i<len(arr1); i++ {
		if arr1[i] > arr2[i] {
			result[i] = arr1[i]
		} else {
			result[i] = arr2[i]
		}
	}
	return result
}

var score int = 0

func visit(x, y, z int) {
	loc := x + y*1000 + z*1000000
	if value, ok := world[loc]; ok {
		if (value == 2) {
			score++
		}
		return
	}
	world[loc] = 1
	if x >= worldMin[0] {
		visit(x-1, y, z)
	}
	if x <= worldMax[0] {
		visit(x+1, y, z)
	}
	if y >= worldMin[1] {
		visit(x, y-1, z)
	}
	if y <= worldMax[1] {
		visit(x, y+1, z)
	}
	if z >= worldMin[2] {
		visit(x, y, z-1)
	}
	if z <= worldMax[2] {
		visit(x, y, z+1)
	}
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

		worldMin = min(xyz, worldMin)
		worldMax = max(xyz, worldMax)
		world[l] = 2
	}

	visit(worldMin[0]-1, worldMin[1], worldMin[2])
	fmt.Println(score)
}
