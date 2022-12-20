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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type H struct {
	value int
	prev  *H
	next  *H
}

func main() {
	var zero *H = nil
	b, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(b), "\n")

	items := make([]*H, 0)

	for _, line := range lines {
		value := strToInt(line)

		item := &H{value, nil, nil}
		if value == 0 {
			zero = item
		}
		items = append(items, item)
	}

	for i, item := range items {
		p := items[(i+len(items)-1)%len(items)]
		item.prev = p
		p2 := items[(i+1)%len(items)]
		item.next = p2
	}

	for _, item := range items {
		if item.value < 0 {
			left(item, abs(item.value))
		}
		if item.value > 0 {
			right(item, abs(item.value))
		}
		fmt.Println("")
	}

	score := 0
	start := zero
	for i := 0; i <= 3000; i++ {
		if i%1000 == 0 && i != 0 {
			fmt.Println("i ", i, start.value)
			score += start.value
		}
		start = start.next
	}

	fmt.Println("score ", score)
}

func extract(item *H) {
	prev := item.prev
	next := item.next

	prev.next = next
	next.prev = prev
}

func addRight(item *H, left *H) {
	right := left.next

	left.next = item
	item.prev = left
	item.next = right
	right.prev = item
}

func left(item *H, n int) {
	for i := 0; i < n; i++ {
		extract(item)
		addRight(item, item.prev.prev)
	}
}

func right(item *H, n int) {
	for i := 0; i < n; i++ {
		extract(item)
		addRight(item, item.next)
	}
}
