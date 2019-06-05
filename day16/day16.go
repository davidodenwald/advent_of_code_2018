package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

const filename = "input.txt"

var b2i = map[bool]int{false: 0, true: 1}
var opts = []func(int, int, int, []int){
	func(a, b, c int, mem []int) { mem[c] = mem[a] + mem[b] },
	func(a, b, c int, mem []int) { mem[c] = mem[a] + b },
	func(a, b, c int, mem []int) { mem[c] = mem[a] * mem[b] },
	func(a, b, c int, mem []int) { mem[c] = mem[a] * b },
	func(a, b, c int, mem []int) { mem[c] = mem[a] & mem[b] },
	func(a, b, c int, mem []int) { mem[c] = mem[a] & b },
	func(a, b, c int, mem []int) { mem[c] = mem[a] | mem[b] },
	func(a, b, c int, mem []int) { mem[c] = mem[a] | b },
	func(a, b, c int, mem []int) { mem[c] = mem[a] },
	func(a, b, c int, mem []int) { mem[c] = a },
	func(a, b, c int, mem []int) { mem[c] = b2i[a > mem[b]] },
	func(a, b, c int, mem []int) { mem[c] = b2i[mem[a] > b] },
	func(a, b, c int, mem []int) { mem[c] = b2i[mem[a] > mem[b]] },
	func(a, b, c int, mem []int) { mem[c] = b2i[a == mem[b]] },
	func(a, b, c int, mem []int) { mem[c] = b2i[mem[a] == b] },
	func(a, b, c int, mem []int) { mem[c] = b2i[mem[a] == mem[b]] },
}

func main() {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("file not found")
	}

	moreThan3 := 0
	optMap := map[int][]int{} // Maps opt number to func index
	seen := map[int]bool{}

	var mem []int
	var cmpMem []int
	var inst []int
	for _, line := range bytes.Split(file, []byte("\n")) {
		if len(line) < 1 {
			po := possibleOpts(mem, cmpMem, inst)
			if len(po) >= 3 {
				moreThan3++
			}

			optCode := inst[0]
			if _, ok := seen[optCode]; !ok {
				seen[optCode] = true
				optMap[optCode] = po
			} else {
				var filtered []int
				for _, opt := range optMap[optCode] {
					if contains(po, opt) {
						filtered = append(filtered, opt)
					}
				}
				optMap[optCode] = filtered
			}
			continue
		}
		switch line[0] {
		case 'B': // line begins with 'Before'
			mem = parseLine(line)
		case 'A': // line begins with 'After'
			cmpMem = parseLine(line)
		default: // line begins with a number
			inst = parseLine(line)
		}
	}
	fmt.Println("Part 1:", moreThan3)

	finMap := map[int]int{}
	for len(optMap) > 0 {
		var filterOpt int
		for k, v := range optMap {
			if len(v) == 1 {
				finMap[k] = v[0]
				filterOpt = v[0]
				delete(optMap, k)
				break
			}
		}
		for k, v := range optMap {
			var filtered []int
			for _, e := range v {
				if e != filterOpt {
					filtered = append(filtered, e)
				}
			}
			optMap[k] = filtered
		}
	}

	prog, err := ioutil.ReadFile("prog.txt")
	if err != nil {
		log.Fatalln("file not found")
	}
	exec(prog, finMap)
}

func parseLine(line []byte) []int {
	var res []int
	numsRe := regexp.MustCompile(`\d+`)
	nums := numsRe.FindAll(line, 4)
	for _, numB := range nums {
		num, _ := strconv.Atoi(string(numB))
		res = append(res, num)
	}
	return res
}

func possibleOpts(startMem, endMem, inst []int) []int {
	var res []int
	origMem := make([]int, len(startMem))
	copy(origMem, startMem)
	for i, opt := range opts {
		opt(inst[1], inst[2], inst[3], startMem)
		if equal(startMem, endMem) {
			res = append(res, i)
		}
		copy(startMem, origMem)
	}
	return res
}

func equal(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func contains(c []int, a int) bool {
	for _, e := range c {
		if e == a {
			return true
		}
	}
	return false
}

func exec(prog []byte, optMap map[int]int) {
	mem := []int{0, 0, 0, 0}
	for _, line := range bytes.Split(prog, []byte("\n")) {
		args := parseLine(line)
		opts[optMap[args[0]]](args[1], args[2], args[3], mem)
	}
	fmt.Println(mem[0])
}
