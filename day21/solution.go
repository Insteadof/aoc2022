package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	value      int
	a          *Monkey
	b          *Monkey
	opperation string
	name       string
}

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

var monkeys map[string]*Monkey = make(map[string]*Monkey, 0)

func main() {
	b, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(b), "\n")

	//root: pppw + sjmn
	for _, line := range lines {
		split := strings.Split(line, ": ")
		value := 0
		if !strings.Contains(split[1], " ") {
			value = strToInt(split[1])
		}

		monkeys[split[0]] = &Monkey{
			value,
			nil,
			nil,
			"",
			split[0],
		}
	}

	for _, line := range lines {
		split := strings.Split(line, ": ")

		if strings.Contains(split[1], " ") {
			j := strings.Split(split[1], " ")
			monkeys[split[0]].opperation = j[1]
			monkeys[split[0]].a = monkeys[j[0]]
			monkeys[split[0]].b = monkeys[j[2]]
		}
	}

	// fmt.Println(get(monkeys["root"].a))
	//fmt.Println(get(monkeys["root"].a, get(monkeys["root"].a,0 ) - get(monkeys["root"].b, 0)))
	var s *int
	if get(monkeys["root"].a, nil) == nil {
		 s = get(monkeys["root"].a, get(monkeys["root"].b, nil))
	} else {
		s = get(monkeys["root"].b, get(monkeys["root"].a, nil))
	}
	fmt.Println(s)
}

func get(monkey *Monkey, goal *int) *int {
	if monkey.name == "humn" {
		if goal != nil {
			fmt.Println("humn", *goal)
		}
		return nil
	}

	if monkey.opperation != "" {
		a := get(monkey.a, nil)
		b := get(monkey.b, nil)
		if goal == nil && (a == nil || b == nil) {
			return nil
		}

		if monkey.opperation == "+" {
			if a == nil {
				h := *goal - *b
				return get(monkey.a, &h)
			}
			if b == nil {
				h := *goal - *a
				return get(monkey.b, &h)
			}
			v := *a + *b
			return &v
		}
		if monkey.opperation == "*" {
			if a == nil {
				h := *goal / *b
				return get(monkey.a, &h)
			}
			if b == nil {
				h := *goal / *a
				return get(monkey.b, &h)
			}
			v := *a * *b
			return &v
		}
		if monkey.opperation == "-" { // a - b = goal
			if a == nil {
				h := *goal + *b
				return get(monkey.a, &h)
			}
			if b == nil {
				h := -(*goal - *a)
				return get(monkey.b, &h)
			}
			v := *a - *b
			return &v
		}
		if monkey.opperation == "/" {  // a / b = goal
			if a == nil {
				h := *goal * *b
				return get(monkey.a, &h)
			}
			if b == nil {
				h := -(*a / *goal)
				return get(monkey.b, &h)
			}
			v := *a / *b
			return &v
		}
	}

	return &monkey.value
}
