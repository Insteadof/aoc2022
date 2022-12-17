package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Valve struct {
	name      string
	rate      int
	open      bool
	tunnels   []*Valve
	beenCount int
}

type ChoiceMap struct {
	choiceI int
	choiceK int
}

var optiomalChoices []ChoiceMap = make([]ChoiceMap, 26)

const ahead = 16

var valves map[string]*Valve = make(map[string]*Valve, 0)
var maxScore int = 0

func main() {
	b, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(b), "\n")

	for _, line := range lines {
		s := strings.Split(line, " ")
		r, _ := strconv.Atoi(strings.TrimSuffix(strings.Split(s[4], "=")[1], ";"))

		valves[s[1]] = &Valve{
			name:      s[1],
			rate:      r,
			tunnels:   make([]*Valve, 0),
			open:      false,
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

	do()

	fmt.Println(maxScore)
}

func do() {
	for level := ahead; level <= 26; level++ {
		fmt.Println("level", level)
		check(valves["AA"], valves["AA"], false, false, 0, level, 0, nil, nil)
	}
}

func check(valve, valve2 *Valve, doOpen1, doOpen2 bool, step int, timeLeft int, score int, cameFrom *Valve, cameFrom2 *Valve) int {
	if doOpen1 {
		valve.open = true
		score += valve.rate * timeLeft
		if score > maxScore {
			maxScore = score
			fmt.Println(score)
		}
	} else {
		valve.beenCount++
	}
	if doOpen2 {
		valve2.open = true
		score += valve2.rate * timeLeft
		if score > maxScore {
			maxScore = score
			fmt.Println(score)
		}
	} else {
		valve2.beenCount++
	}

	localMax := score

	if timeLeft > 1 {
		for i := 0; i <= len(valve.tunnels); i++ {
			tunnel := valve
			if i == 0 {
				if valve.rate == 0 || valve.open {
					continue
				}
			} else {
				tunnel = valve.tunnels[i-1]

				if tunnel.beenCount >= len(tunnel.tunnels) {
					continue
				}
				// if (tunnel.rate == 0) && (cameFrom == tunnel || cameFrom2 == tunnel) {
				// 	continue
				// }
				if (valve.rate == 0) && (cameFrom == tunnel) {
					continue
				}
			}

			for k := 0; k <= len(valve2.tunnels); k++ {
				tunnel2 := valve2
				if k == 0 {
					if valve2.rate == 0 || valve2.open || valve == valve2 {
						continue
					}
				} else {
					if i != 0 && timeLeft <= 2 {
						continue
					}
					tunnel2 = valve2.tunnels[k-1]
					if tunnel2.beenCount >= len(tunnel2.tunnels) {
						continue
					}
					// if (tunnel2.rate == 0) && (cameFrom == tunnel2 || cameFrom2 == tunnel2) {
					// 	continue
					// }
					if (valve2.rate == 0) && (cameFrom2 == tunnel2) {
						continue
					}
				}

				if timeLeft > ahead {
					if optiomalChoices[step].choiceI != i || optiomalChoices[step].choiceK != k {
						continue
					}
				}

				subScore := check(tunnel, tunnel2, i == 0, k == 0, step+1, timeLeft-1, score, valve, valve2)
				if subScore > localMax {
					localMax = subScore
					optiomalChoices[step].choiceI = i
					optiomalChoices[step].choiceK = k
				}
			}
		}
	}

	if doOpen1 {
		valve.open = false
	} else {
		valve.beenCount--
	}

	if doOpen2 {
		valve2.open = false
	} else {
		valve2.beenCount--
	}

	return localMax
}
