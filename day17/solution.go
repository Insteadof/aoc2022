package main

import (
	"fmt"
	"os"
)

var world map[int64][]int = make(map[int64][]int, 0)

var rockData = [][][]int{
	{
		{1, 1, 1, 1},
	},
	{
		{0, 1, 0},
		{1, 1, 1},
		{0, 1, 0},
	},
	{
		{1, 1, 1},
		{0, 0, 1},
		{0, 0, 1},
	},
	{
		{1},
		{1},
		{1},
		{1},
	},
	{
		{1, 1},
		{1, 1},
	},
}

var rockX int = 0
var rockY int64 = 0
var rockType int = -1
var extraWorld int64 = 0

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {
	b, _ := os.ReadFile("input.txt")

	line := string(b)
	fmt.Println(len(line))
	pos := 0

	for i := 0; i < 200; i++ {
		dropRock()
		if pos == 0 {
			fmt.Println("reset", len(world))
		}
		count := 0
		for fall(line[pos : pos+1], count) {
			count++
			pos = (pos + 1) % len(line)
		}
		pos = (pos + 1) % len(line)

	  // fmt.Println("world", rockType, rockX, pos)
		if (rockType==0 && rockX == 1 && pos==2921) {
			fmt.Println("reset1", i, worldMaxY())
		}
		if (rockType==1 && rockX == 0 && pos==2925) {
			fmt.Println("reset2", i, worldMaxY())
		}
		if (rockType==2 && rockX == 1 && pos==2929) {
			fmt.Println("reset3", i, worldMaxY())
		}
		if (rockType==3 && rockX == 2 && pos==2935) {
			fmt.Println("reset4", i, worldMaxY())
		}
		if (rockType==4 && rockX == 2 && pos==2939) {
			fmt.Println("reset5", i, worldMaxY())
		}

		// every 1700 i, pattern repeats
		// 2623 height increase
		// print(1000000000000 % 1700)
		// print((1000000000000 // 1700) * 2623 + 318) //318=200 iterations

		for yy, row := range rockData[rockType] {
			for xx, cell := range row {
				if cell == 1 {
					settle(rockType+1, rockX+xx, rockY+int64(yy))
				}
			}
		}
	}
	dropRock()
}

func collision(x int, y int64) bool {
	if x < 0 || x >= 7 || y < 0 {
		return true
	}
	row, ok := world[y]
	if !ok {
		return false
	}
	return row[x] > 0
}

func settle(n, x int, y int64) {
	_, ok := world[y]
	if !ok {
		world[y] = []int{0, 0, 0, 0, 0, 0, 0}
	}
	world[y][x] = n
}

func rockCollision(_type, x int, y int64) bool {
	for yy, row := range rockData[_type] {
		for xx, cell := range row {
			if cell == 1 {
				if collision(x+xx, y+int64(yy)) {
					return true
				}
			}
		}
	}
	return false
}

func fall(push string, age int) bool {
	if push == ">" {
		if (age <= 2) {
			if (rockX + len(rockData[rockType][0]) < 7) {
				rockX++
			}
		} else {
			if !rockCollision(rockType, rockX+1, rockY) {
				rockX++
			}
		}
	} else if push == "<" {
		if (age <= 2) {
			if (rockX > 0) {
				rockX--
			}
		} else {
			if !rockCollision(rockType, rockX-1, rockY) {
				rockX--
			}
		}
	} else {
		fmt.Println("Unknown push", push)
	}
	if (age <= 2) {
		rockY--
		return true
	}

	if !rockCollision(rockType, rockX, rockY-1) {
		rockY--
		return true
	} else {
		return false
	}
}

func worldMaxY() int64 {
	return extraWorld + int64(len(world))
}

func dropRock() {
	if len(world) > 0 {
		if len(world) > 1000 {
			newWorld := make(map[int64][]int, 100)

			max := extraWorld + int64(len(world))
			for i:=(max-1);i > max-100;i-- {
				newWorld[i] = world[i]
			}
			extraWorld += int64(len(world)-len(newWorld))
			world = newWorld
		}

		fmt.Println("world hight ", worldMaxY())
		rockY = worldMaxY() + 3
	} else {
		rockY = 3
	}

	rockX = 2
	rockType = (rockType + 1) % len(rockData)
}
