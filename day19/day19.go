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
var opts = map[string]func(int, int, int, []int){
	"addr": func(a, b, c int, mem []int) { mem[c] = mem[a] + mem[b] },
	"addi": func(a, b, c int, mem []int) { mem[c] = mem[a] + b },
	"mulr": func(a, b, c int, mem []int) { mem[c] = mem[a] * mem[b] },
	"muli": func(a, b, c int, mem []int) { mem[c] = mem[a] * b },
	"banr": func(a, b, c int, mem []int) { mem[c] = mem[a] & mem[b] },
	"bani": func(a, b, c int, mem []int) { mem[c] = mem[a] & b },
	"borr": func(a, b, c int, mem []int) { mem[c] = mem[a] | mem[b] },
	"bori": func(a, b, c int, mem []int) { mem[c] = mem[a] | b },
	"setr": func(a, b, c int, mem []int) { mem[c] = mem[a] },
	"seti": func(a, b, c int, mem []int) { mem[c] = a },
	"gtir": func(a, b, c int, mem []int) { mem[c] = b2i[a > mem[b]] },
	"gtri": func(a, b, c int, mem []int) { mem[c] = b2i[mem[a] > b] },
	"gtrr": func(a, b, c int, mem []int) { mem[c] = b2i[mem[a] > mem[b]] },
	"eqir": func(a, b, c int, mem []int) { mem[c] = b2i[a == mem[b]] },
	"eqri": func(a, b, c int, mem []int) { mem[c] = b2i[mem[a] == b] },
	"eqrr": func(a, b, c int, mem []int) { mem[c] = b2i[mem[a] == mem[b]] },
}

type instruction struct {
	name string
	args []int
}

func main() {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("file not found")
	}
	ip, insts := parseInput(file)
	mem := make([]int, 6)
	fmt.Println("Part 1:", calcRes(ip, insts, mem))
	mem = []int{1, 0, 0, 0, 0, 0}
	fmt.Println("Part 2:", calcRes(ip, insts, mem))
}

func parseInput(in []byte) (int, []instruction) {
	numRe := regexp.MustCompile(`\d+`)
	lines := bytes.Split(in, []byte("\n"))
	ip, _ := strconv.Atoi(string(numRe.Find(lines[0])))

	var ins []instruction
	for _, line := range lines[1:] {
		name := string(line[:4])
		args := []int{}
		for _, b := range numRe.FindAll(line, 3) {
			num, _ := strconv.Atoi(string(b))
			args = append(args, num)
		}
		ins = append(ins, instruction{
			name: name,
			args: args,
		})
	}
	return ip, ins
}

func calcRes(ip int, insts []instruction, mem []int) int {
	i := 0
	for mem[ip] < len(insts) {
		ins := insts[mem[ip]]
		opts[ins.name](ins.args[0], ins.args[1], ins.args[2], mem)

		// break out loop when mem[2] has been initialized
		if mem[ip] == 0 && i > 0 {
			break
		}
		mem[ip]++
		i++
	}
	return sumFactors(mem[2])
}

func sumFactors(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}
