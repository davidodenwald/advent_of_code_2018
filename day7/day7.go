package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

const filename = "example.txt"

func main() {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("file not found")
	}
	part1(string(file))
}

func part1(data string) {
	steps := make(map[byte]int)
	depend := make(map[byte][]byte)
	for _, line := range strings.Split(data, "\n") {
		step := line[5]
		after := line[36]
		steps[after]++
		depend[step] = append(depend[step], after)
	}
	var start []byte
	for step := range depend {
		if steps[step] == 0 {
			start = append(start, step)
		}
	}

	var res string
	for len(start) > 0 {
		tmpStart := make([]byte, len(start))
		copy(tmpStart, start)
		sort.Slice(tmpStart, func(i, j int) bool {
			return tmpStart[i] < tmpStart[j]
		})
		x := tmpStart[0]

		start = start[:0]
		for _, y := range tmpStart {
			if y != x {
				start = append(start, y)
			}
		}

		res += string(x)
		for _, y := range depend[x] {
			steps[y]--
			if steps[y] == 0 {
				start = append(start, y)
			}
		}
	}
	fmt.Println(res)
}
