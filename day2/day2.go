package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const filename = "input.txt"

func main() {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("file not found")
	}
	lines := strings.Split(string(file), "\n")

	part1(lines)
	part2(lines)
}

func part1(ids []string) {
	var twos int
	var threes int

	for _, id := range ids {
		chars := make(map[rune]int)
		for _, char := range id {
			chars[char]++
		}

		addedTwo := false
		addedThree := false

		for _, num := range chars {
			switch num {
			case 2:
				if !addedTwo {
					twos++
					addedTwo = true
				}
			case 3:
				if !addedThree {
					threes++
					addedThree = true
				}
			}
		}
	}
	fmt.Println(twos * threes)
}

func part2(ids []string) {
	for _, id := range ids {
		for _, cmpID := range ids {
			if cmpID == id {
				continue
			}
			for i := 0; i < len(id) && i < len(cmpID); i++ {
				if id[:i]+id[i+1:len(id)] == cmpID[:i]+cmpID[i+1:len(cmpID)] {
					fmt.Println(id[:i] + id[i+1:len(id)])
					return
				}
			}
		}
	}
}
