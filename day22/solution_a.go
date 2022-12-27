package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

var world [][]int = make([][]int, 0)
var width = 0
var height = 0

func main() {
	b, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(b), "\n\n")
	grid := lines[0]
	gridLines := strings.Split(grid, "\n")
	path := lines[1]

	width = len(gridLines[0])
	height = len(gridLines)

	for _, line := range gridLines {
		yy := make([]int, width, width)
		for x, char := range line {
			v := 0
			if char == '#' {
				v = 2
			}
			if char == '.' {
				v = 1
			}
			yy[x] = v
		}
		world = append(world, yy)
	}

	turns := strings.FieldsFunc(path, func(r rune) bool {
		return r >= '0' && r <= '9'
	})
	amounts := strings.FieldsFunc(path, func(r rune) bool {
		return (r == 'L' || r == 'R')
	})
	x, y, dir := walk(turns, amounts)
	fmt.Println(x, y, dir, (y+1) * 1000 + (x+1) * 4 + dir)
}

func walk(turns []string, amounts []string) (x int, y int, dir int) {
	x = 0
	y = 0
	for world[y][x] != 1 {
		x++
	}
	dir = 0
	for i := 0; i < len(amounts); i++ {
		amount := strToInt(amounts[i])
		x, y = steps(x, y, dir, amount)
		fmt.Println(x, y, dir)

		if i < len(turns) {
			if turns[i] == "L" {
				dir = (dir + 3) % 4
			}
			if turns[i] == "R" {
				dir = (dir + 1) % 4
			}
		}
	}
	return
}

func steps(x, y, dir, n int) (xx int, yy int) {
	xx = x
	yy = y
	for i := 0; i < n; i++ {
		var success bool
		xx, yy, success = step(xx, yy, dir)
		if (!success) {
			return
		}
	}
	return
}

func step(x, y, dir int) (xx int, yy int, success bool) {
	success = true
	xx = x
	yy = y
	switch dir {
	case 0:
		xx = (x + 1) % width
	case 1:
		yy = (y + 1) % height
	case 2:
		xx = (x - 1 + width) % width
	case 3:
		yy = (y - 1 + height) % height
	}

	if world[yy][xx] == 2 {
		return x, y, false
	}
	if world[yy][xx] == 0 {
		xx, yy, s := step(xx, yy, dir)
		if !s {
			return x, y, false
		}
		return xx, yy, true
	}
	return
}
