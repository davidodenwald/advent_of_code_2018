package main

import (
	"bytes"
	"fmt"
)

func main() {
	input := []byte{3, 7}
	before := 864801
	res := part1(input, before, 10)

	for _, x := range res {
		fmt.Print(x)
	}
	fmt.Println()

	fmt.Println(part2(input, []byte{8, 6, 4, 8, 0, 1}))
}

func part1(input []byte, before, recipes int) []byte {
	e1 := 0
	e2 := 1

	for i := 0; i < recipes+before; i++ {
		sum := input[e1] + input[e2]
		if sum > 9 {
			input = append(input, sum/10)
			input = append(input, sum%10)
		} else {
			input = append(input, sum)
		}

		e1 += int(input[e1]) + 1
		e1 = e1 % len(input)
		e2 += int(input[e2]) + 1
		e2 = e2 % len(input)
	}
	return input[before : before+recipes]
}

func part2(input, score []byte) int {
	index := -1
	e1 := 0
	e2 := 1

	for i := 0; index < 0; i++ {
		sum := input[e1] + input[e2]
		if sum > 9 {
			input = append(input, sum/10)
			input = append(input, sum%10)
		} else {
			input = append(input, sum)
		}

		e1 += int(input[e1]) + 1
		e1 = e1 % len(input)
		e2 += int(input[e2]) + 1
		e2 = e2 % len(input)

		if i%100000 == 0 {
			index = bytes.Index(input, score)
		}
	}
	return index
}
