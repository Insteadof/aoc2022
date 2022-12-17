package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Valve struct {
	name    string
	rate    uint
	open    bool
	tunnels []*Valve
	beenCount int
}

var valves map[string]*Valve = make(map[string]*Valve, 0)
var maxScore uint = 0

func main() {
	b, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(b), "\n")

	for _, line := range lines {
		s := strings.Split(line, " ")
		r, _ := strconv.Atoi(strings.TrimSuffix(strings.Split(s[4], "=")[1], ";"))

		valves[s[1]] = &Valve{
			name:    s[1],
			rate:    uint(r),
			tunnels: make([]*Valve, 0),
			open:    false,
			beenCount: 0,
		}
	}

	for _, line := range lines {
		s := strings.Split(line, " ")
		valve := valves[s[1]]

		for i := 9; i < len(s); i++ {
			v := strings.TrimSuffix(s[i], ",")
			valve.tunnels = append(valve.tunnels, valves[v])
		}
	}

	fmt.Println(valves["HH"])

	check(valves["AA"], 30, 0, false, 0, nil)

	fmt.Println(maxScore)
}

func check(valve *Valve, timeLeft uint, score uint, canOpen bool, zeros int, cameFrom *Valve) {
	valve.beenCount++

	didOpen := false
	if canOpen && valve.rate > 0 && !valve.open && timeLeft > 1 {
		didOpen = true
		valve.open = true
		timeLeft--
		score += valve.rate * timeLeft
		if score > maxScore {
			maxScore = score
			fmt.Println(score)
		}
	}

	if valve.rate == 0 {
		zeros++
	}

	if timeLeft > 2 {
		for _, tunnel := range valve.tunnels {
			if (tunnel.beenCount < len(tunnel.tunnels)) {
				if (tunnel.rate == 0 && (zeros > 16) || (cameFrom == tunnel)) {
					continue
				}

				if tunnel.rate > 0 && !tunnel.open {
					check(tunnel, timeLeft-1, score, true, zeros, valve)
				}

				if tunnel.rate < 10 || tunnel.open {
					check(tunnel, timeLeft-1, score, false, zeros, valve)
				}
			}
		}
	}

	if didOpen {
		valve.open = false
	}

	valve.beenCount--
}
