package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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
	rand.Seed(time.Now().UnixNano())

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

	xx, yy, _, _ := step(50, 81, 2)
	x, y, _, _ := step(xx, yy, 3)

	for i := 0; i < 10000000; i++ {
		x := rand.Intn(150)
		y := rand.Intn(200)
		dir := rand.Intn(4)
		if dir == 4 {
			panic("dir = 4")
		}
		if world[y][x] == 1 {
			xx, yy, dirOut, succes := step(x, y, dir)
			if succes {
				xx, yy, dirOut, succes = step(xx, yy, (dirOut+2)%4)
				if !succes {
					panic("no succes")
				}
				if xx != x || yy != y || dirOut != (dir+2)%4 {
					panic(fmt.Sprintf("not back %x,%x", x, y))
				}
			}
		}
	}

	turns := strings.FieldsFunc(path, func(r rune) bool {
		return r >= '0' && r <= '9'
	})
	amounts := strings.FieldsFunc(path, func(r rune) bool {
		return (r == 'L' || r == 'R')
	})
	x, y, dir := walk(turns, amounts)
	fmt.Println(x, y, dir, (y+1)*1000+(x+1)*4+dir)
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
		x, y, dir = steps(x, y, dir, amount)
		// fmt.Println(x, y, dir)

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

func steps(x, y, dir, n int) (xx int, yy int, dirOut int) {
	xx = x
	yy = y
	for i := 0; i < n; i++ {
		var success bool
		xx, yy, dir, success = step(xx, yy, dir)
		dirOut = dir
		if !success {
			return
		}
	}
	return
}

func step(x, y, dir int) (xx int, yy int, dirOut int, success bool) {
	success = true
	xx = x
	yy = y
	reset := func(xxx, yyy int, dirOut int, s bool) (int, int, int, bool) {
		if !s {
			return x, y, dir, s
		}
		return xxx, yyy, dirOut, s
	}

	if y >= 0 && y < 50 {
		switch dir {
		case 0:
			if x == 149 {
				return reset(step(100, 100+(49-y), 2))
			}
		case 1:
			if x >= 100 && y == 49 {
				return reset(step(100, x-50, 2))
			}
		case 2:
			if x == 50 {
				return reset(step(-1, 100+(49-y), 0))
			}
		case 3:
			if y == 0 {
				if x >= 100 {
					return reset(step(x-100, 200, 3))
				}
				return reset(step(-1, x+100, 0))
			}
		}
	} else if y >= 50 && y < 100 {
		switch dir {
		case 0:
			if x == 99 {
				return reset(step(y+50, 50, 3))
			}
		case 2:
			if x == 50 {
				return reset(step(y - 50, 99, 1))
			}
		}
	} else if y >= 100 && y < 150 {
		switch dir {
		case 0:
			if x == 99 {
				return reset(step(150, (149 - y), 2))
			}
		case 1:
			if x >= 50 && y == 149 {
				return reset(step(50, 100+x, 2))
			}
		case 2:
			if x == 0 {
				return reset(step(49, 100+(49-y), 0))
			}
		case 3:
			if x < 50 && y == 100 {
				return reset(step(49, x+50, 0))
			}
		}
	} else if y >= 150 && y < 200 {
		switch dir {
		case 0:
			if x == 49 {
				return reset(step(y-100, 150, 3))
			}
		case 1:
			if y == 199 {
				return reset(step(x+100, -1, 1))
			}
		case 2:
			if x == 0 {
				return reset(step(y-100, -1, 1))
			}
		}
	} else {
		// fmt.Println("error")
	}
	dirOut = dir

	switch dir {
	case 0:
		xx = x + 1
	case 1:
		yy = y + 1
	case 2:
		xx = x - 1
	case 3:
		yy = y - 1
	}

	if world[yy][xx] == 2 {
		return x, y, dir, false
	}
	if world[yy][xx] == 0 {
		fmt.Print("error")
		// xx, yy, s := step(xx, yy, dir)
		// if !s {
		// 	return x, y, false
		// }
		// return xx, yy, true
	}
	return
}
