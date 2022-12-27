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

type Elve struct {
	dir   int
	x     int
	y     int
	nextX int
	nextY int
	block bool
	track bool
}

var offset = 200

var world [][]*Elve = make([][]*Elve, 500)
var elves []*Elve = make([]*Elve, 0)

var minX = 1000000
var maxX = 0
var minY = 10000000
var maxY = 0

func main() {
	for i := 0; i < 500; i++ {
		world[i] = make([]*Elve, 500)
	}

	b, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(b), "\n")

	for y, line := range lines {
		fmt.Println(line)
		for x, c := range line {
			if c == '#' {
				xx := x + offset
				yy := y + offset
				world[yy][xx] = &Elve{
					dir:   0,
					x:     xx,
					y:     yy,
					block: false,
					nextX: 0,
					nextY: 0,
					track: x == 3 && y == 4,
				}
				updateMinMax(xx, yy)
				elves = append(elves, world[yy][xx])
			}
		}
	}

	for i := 0; i < 1000000; i++ {
		if !round() {
			fmt.Println("did not move after round", i+1)
			break
		}
		fmt.Println("after round", i+1)
		for y := minY; y <= maxY; y++ {
			line := ""
			for x := minX; x <= maxX; x++ {
				if world[y][x] == nil {
					line += "."
				} else {
					line += "#"
				}
			}
			fmt.Println(line)
		}
	}

	score := 0
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if world[y][x] == nil {
				score++
			}
		}
	}

	fmt.Println("score", score)
}

func round() (didMove bool) {
	didMove = false
	for _, elve := range elves {
		otherElve := false
		move := -1
		for i := 0; i < 4; i++ {
			dir := (elve.dir + i) % 4
			x := elve.x
			y := elve.y

			if dir == 0 {
				if world[y-1][x] != nil || world[y-1][x-1] != nil || world[y-1][x+1] != nil {
					otherElve = true
				} else {
					if move == -1 {
						elve.nextX = x
						elve.nextY = y - 1
						move = 0
					}
				}
			}
			if dir == 1 {
				if world[y+1][x] != nil || world[y+1][x-1] != nil || world[y+1][x+1] != nil {
					otherElve = true
				} else {
					if move == -1 {
						elve.nextX = x
						elve.nextY = y + 1
						move = 1
					}
				}
			}
			if dir == 2 {
				if world[y-1][x-1] != nil || world[y][x-1] != nil || world[y+1][x-1] != nil {
					otherElve = true
				} else {
					if move == -1 {
						elve.nextX = x - 1
						elve.nextY = y
						move = 2
					}
				}
			}
			if dir == 3 {
				if world[y-1][x+1] != nil || world[y][x+1] != nil || world[y+1][x+1] != nil {
					otherElve = true
				} else {
					if move == -1 {
						elve.nextX = x + 1
						elve.nextY = y
						move = 3
					}
				}
			}
		}

		elve.dir++
		if elve.track {
			fmt.Println("")
		}

		if !otherElve {
			elve.nextX = 0
			elve.nextY = 0

			continue
		}

		//if move != -1 {
		//= (move + 1) % 4
		//}
	}

	for _, elve := range elves {
		if elve.nextX == 0 {
			continue
		}

		if world[elve.nextY][elve.nextX] != nil {
			elve2 := world[elve.nextY][elve.nextX]
			elve.block = true
			elve2.block = true
			elve.nextX = 0
			elve.nextY = 0
		} else {
			world[elve.nextY][elve.nextX] = elve
			world[elve.y][elve.x] = nil
		}
	}

	for _, elve := range elves {
		if elve.block && elve.nextX != 0 {
			world[elve.nextY][elve.nextX] = nil
			world[elve.y][elve.x] = elve
			elve.nextX = 0
			elve.nextY = 0
		}
		elve.block = false

		if elve.nextX == 0 {
			continue
		}

		elve.x = elve.nextX
		elve.y = elve.nextY
		didMove = true
		updateMinMax(elve.x, elve.y)
	}
	return
}

func updateMinMax(x, y int) {
	if x < minX {
		minX = x
	}
	if x > maxX {
		maxX = x
	}
	if y < minY {
		minY = y
	}
	if y > maxY {
		maxY = y
	}
}
