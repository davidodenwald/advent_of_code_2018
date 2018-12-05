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
	part1(string(file))
	part2(string(file))
}

func toLower(b byte) byte {
	return b + 32
}

func collapsePoly(poly string) string {
	changed := true

	for changed {
		changed = false
		for i := 0; i < len(poly)-1; i++ {
			if toLower(poly[i]) == poly[i+1] || poly[i] == toLower(poly[i+1]) {
				poly = poly[:i] + poly[i+2:]
				changed = true
			}
		}
	}
	return poly
}

func part1(poly string) {
	poly = collapsePoly(poly)
	fmt.Println(len(poly))
}

func part2(poly string) {
	chars := make(map[byte]bool)

	for i := 0; i < len(poly); i++ {
		if poly[i] < 'a' {
			chars[toLower(poly[i])] = true
		} else {
			chars[poly[i]] = true
		}
	}

	var newPoly string
	min := 1<<63 - 1
	for char := range chars {
		newPoly = strings.Replace(poly, strings.ToUpper(string(char)), "", -1)
		newPoly = strings.Replace(newPoly, string(char), "", -1)
		curLen := len(collapsePoly(newPoly))
		if curLen < min {
			min = curLen
		}
	}
	fmt.Println(min)
}
