package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cost map[string]int

type Blueprint struct {
	id       int
	robot    map[string]Cost
	maxCosts map[string]int
}

type Resources struct {
	product map[string]int
	robots  map[string]int
}

func strToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func indexOf(data []string, element string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func getCost(line []string, robotName string) int {
	robotIndex := indexOf(line, robotName)
	if robotIndex == -1 {
		return 0
	}
	return strToInt(line[robotIndex-1])
}

func main() {
	b, _ := os.ReadFile("input.txt")

	bluePrints := make([]Blueprint, 0)
	//Blueprint 1: Each ore robot costs 3 ore. Each clay robot costs 4 ore. Each obsidian robot costs 2 ore and 20 clay. Each geode robot costs 4 ore and 7 obsidian.
	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		line = strings.Replace(line, "Blueprint ", "", 1)
		g := strings.Split(line, ":")

		robots := strings.Split(g[1], ".")

		bluePrint := Blueprint{
			id:    strToInt(g[0]),
			robot: make(map[string]Cost),
			maxCosts: map[string]int{
				"ore":      0,
				"clay":     0,
				"obsidian": 0,
				"geode":    1000,
			},
		}
		for _, robot := range robots {
			if robot == "" {
				continue
			}
			r := strings.Split(robot, " ")

			name := r[indexOf(r, "robot")-1]
			r = r[indexOf(r, "robot"):]
			cost := Cost{
				"ore":      getCost(r, "ore"),
				"clay":     getCost(r, "clay"),
				"obsidian": getCost(r, "obsidian"),
				"geode":    getCost(r, "geode"),
			}
			for key, value := range cost {
				bluePrint.maxCosts[key] = Max(bluePrint.maxCosts[key], value)
			}
			bluePrint.robot[name] = cost
		}

		bluePrints = append(bluePrints, bluePrint)
	}

	score := 0
	for i, print := range bluePrints {
		score += print.id * algo(print, Resources{
			product: map[string]int{
				"ore":      0,
				"clay":     0,
				"obsidian": 0,
				"geode":    0,
			},
			robots: map[string]int{
				"ore":      1,
				"clay":     0,
				"obsidian": 0,
				"geode":    0,
			},
		}, 24)
		fmt.Println("score: ", score, "blueprint: ", i+1)
	}

	fmt.Println("score: ", score)
}

var robots = []string{"geode", "obsidian", "ore", "clay"}

func algo(bluePrint Blueprint, r Resources, timeLeft int) int {
	if timeLeft == 1 {
		bla := r.product["geode"] + r.robots["geode"]
		return bla
	}

	run := robots
	options := []string{""}
	if timeLeft <= 3 {
		run = []string{"geode"}
	}

	for _, robot := range run {
		if r.robots[robot] >= bluePrint.maxCosts[robot] {
			continue
		}
		if timeLeft < 11 && robot == "ore" { // && r.robots[robot] >= bluePrint.maxCosts[robot] / 2) {
			continue
		}
		if timeLeft < 8 && robot == "clay" { // && r.robots[robot] >= bluePrint.maxCosts[robot] / 2) {
			continue
		}

		gotall := true
		for material, amount := range bluePrint.robot[robot] {
			if r.product[material] < amount {
				gotall = false
				break
			}
		}

		if gotall {
			if robot == "geode" {
				options = []string{robot}
				break
			} else {
				options = append(options, robot)
			}
		}
	}

	max := 0

	for _, option := range options {
		newResources := Resources{
			product: make(map[string]int, 4),
			robots:  make(map[string]int, 4),
		}
		for key, v := range r.robots {
			newResources.robots[key] = v
		}
		for k, v := range r.product {
			newResources.product[k] = v + r.robots[k]
		}

		if option != "" {
			// create robot
			newResources.robots[option] += 1
			for material, amount := range bluePrint.robot[option] {
				newResources.product[material] -= amount
			}
		}

		localMax := algo(bluePrint, newResources, timeLeft-1)
		if localMax > max {
			max = localMax
		}
	}
	return max
}
