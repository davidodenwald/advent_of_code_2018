package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	filename    = "input.txt"
	generations = 20
)

func main() {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("file not found")
	}
	state, rules := parseInput(string(file))

	plantIndex := 0
	for gen := 0; gen < generations; gen++ {
		firstPlant := len(state)
		lastPlant := 0
		for i := 0; i < len(state); i++ {
			if state[i] == '#' && i < firstPlant {
				firstPlant = i
			}
			if state[i] == '#' && i > lastPlant {
				lastPlant = i
			}
		}

		if 4-(len(state)-lastPlant) > 0 {
			pots := make([]rune, 4-(len(state)-lastPlant))
			for i := range pots {
				pots[i] = '.'
			}
			state = append(state, pots...)
		}

		if firstPlant < 5 {
			pots := make([]rune, 5-firstPlant)
			for i := range pots {
				pots[i] = '.'
				plantIndex++
			}
			state = append(pots, state...)
		}

		newState := make([]rune, len(state))
		copy(newState, state)
		for i := 3; i < len(state)-2; i++ {
			area := state[i-2 : i+3]
			if rules[string(area)] != '#' {
				newState[i] = '.'
			} else {
				newState[i] = rules[string(area)]
			}
		}
		state = newState
	}

	count := 0
	for i, r := range state {
		if r == '#' {
			count += i - plantIndex
		}
	}
	fmt.Println(count)
}

func parseInput(input string) ([]rune, map[string]rune) {
	rules := make(map[string]rune)
	lines := strings.Split(input, "\n")
	state := strings.TrimLeft(lines[0], "initial state: ")

	for i := 2; i < len(lines); i++ {
		rule := strings.Split(lines[i], " => ")
		rules[rule[0]] = []rune(rule[1])[0]
	}

	return []rune(state), rules
}
