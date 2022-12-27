package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Wind struct {
	dir int
	x   int
	y   int
}

type Pos struct {
	x int
	y int
}

var winds = make([]*Wind, 0)
var running = true

var world [][]int
var width int
var height int

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

	world = make([][]int, 0)
	lines := strings.Split(string(b), "\n")
	for i := 1; i < len(lines)-1; i++ {
		worldY := make([]int, len(lines[0])-2)
		for x, c := range lines[i] {
			if c != '#' {
				if c != '.' {
					dir := 0
					switch c {
					case '>':
						dir = 1
					case 'v':
						dir = 2
					case '<':
						dir = 3
					}
					wind := Wind{
						dir: dir,
						x:   x - 1,
						y:   i - 1,
					}
					winds = append(winds, &wind)
					worldY[x-1] = 1
				} else {
					worldY[x-1] = 0
				}
			}
		}
		world = append(world, worldY)
	}
	width = len(world[0])
	height = len(world)

	steps := solve([]Pos{{0, -1}}, Pos{0, -1}, Pos{width - 1, height}, 1)
	fmt.Println(steps)
	running = true
	steps += solve([]Pos{{width - 1, height}}, Pos{width - 1, height}, Pos{0, -1}, 1)
	running = true
	steps += solve([]Pos{{0, -1}}, Pos{0, -1}, Pos{width - 1, height}, 1)
	fmt.Println("total", steps)
}

func updateWinds() {
	for _, wind := range winds {
		world[wind.y][wind.x]--
		switch wind.dir {
		case 0:
			wind.y = (wind.y + height - 1) % height
		case 1:
			wind.x = (wind.x + 1) % width
		case 2:
			wind.y = (wind.y + 1) % height
		case 3:
			wind.x = (wind.x + width - 1) % width
		}
		world[wind.y][wind.x]++
	}
}

func appendIfPossible(x, y int, pos []Pos, start, goal Pos, n int) []Pos {
	if x == start.x && y == start.y {
		for _, p := range pos {
			if p.x == x && p.y == y {
				return pos
			}
		}
		pos = append(pos, Pos{x, y})
	}
	if x == goal.x && y == goal.y {
		fmt.Println("Found it!", n)
		running = false
	}

	if x < 0 || y < 0 || x >= width || y >= height {
		return pos
	}

	if world[y][x] == 0 {
		for _, p := range pos {
			if p.x == x && p.y == y {
				return pos
			}
		}
		pos = append(pos, Pos{x, y})
	}
	return pos
}

func solve(current []Pos, start, goal Pos, n int) int {
	next := make([]Pos, 0)

	updateWinds()

	for _, pos := range current {
		next = appendIfPossible(pos.x, pos.y, next, start, goal, n)
		next = appendIfPossible(pos.x, pos.y+1, next, start, goal, n)
		next = appendIfPossible(pos.x, pos.y-1, next, start, goal, n)
		next = appendIfPossible(pos.x+1, pos.y, next, start, goal, n)
		next = appendIfPossible(pos.x-1, pos.y, next, start, goal, n)
	}
	// fmt.Println("space", len(next), n)
	//fmt.Println("", next, n)
	if running {
		return solve(next, start, goal, n+1)
	} else {
		return n
	}
}
