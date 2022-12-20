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

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

type H struct {
	value int64
	prev  *H
	next  *H
}

var listLength int64 = 0

func main() {
	var zero *H = nil
	b, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(b), "\n")

	items := make([]*H, 0)

	for _, line := range lines {
		value := strToInt(line)

		item := &H{int64(value * 811589153), nil, nil}
		if value == 0 {
			zero = item
		}
		items = append(items, item)
	}

	listLength = int64(len(items))

	for i, item := range items {
		item.prev = items[(i+len(items)-1)%len(items)]
		item.next = items[(i+1)%len(items)]
	}

	for k := 0; k < 10; k++ {
		for _, item := range items {
			if item.value < 0 {
				left(item, abs(item.value))
			}
			if item.value > 0 {
				right(item, abs(item.value))
			}
		}
		fmt.Println("")
	}

	score := int64(0)
	start := zero
	for i := 0; i <= 3000; i++ {
		if i%1000 == 0 && i != 0 {
			fmt.Println("i ", i, start.value * 1)
			score += start.value * 1
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

func insertRight(item *H, left *H) {
	right := left.next

	left.next = item
	item.prev = left
	item.next = right
	right.prev = item
}

func left(item *H, n int64) {
	for i := 0; i < int(n % (listLength-1)); i++ {
		extract(item)
		insertRight(item, item.prev.prev)
	}
}

func right(item *H, n int64) {
	for i := 0; i < int(n % (listLength-1)); i++ {
		extract(item)
		insertRight(item, item.next)
	}
}
