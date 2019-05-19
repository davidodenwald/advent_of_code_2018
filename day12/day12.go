package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	filename    = "example.txt"
	generations = 20
)

func main() {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("file not found")
	}
	state, rules := parseInput(string(file))

	for gen := 0; gen < generations; gen++ {
		fmt.Println(string(state))
		for i := 3; i < len(state)-2; i++ {
			area := state[i-3 : i+2]
			if rules[string(area)] != '#' {
				state[i] = '.'
			} else {
				state[i] = rules[string(area)]
			}
		}
	}

	// for k, v := range rules {
	// 	fmt.Printf("%s -> %c\n", k, v)
	// }
}

func parseInput(input string) ([]rune, map[string]rune) {
	rules := make(map[string]rune)
	lines := strings.Split(input, "\n")
	state := strings.TrimLeft(lines[0], "initial state: ")
	state = ".." + state + ".."

	for i := 2; i < len(lines); i++ {
		rule := strings.Split(lines[i], " => ")
		rules[rule[0]] = []rune(rule[1])[0]
	}

	return []rune(state), rules
}
