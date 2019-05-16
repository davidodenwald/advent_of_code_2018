package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"

	"github.com/golang-collections/collections/stack"
)

const filename = "input.txt"

type node struct {
	childLen int
	entryLen int
	Childs   []*node
	Entries  []int
}

func (n node) String() string {
	return fmt.Sprintf("Childs: %v Entries: %v", n.childLen, n.Entries)
}

func main() {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("file not found")
	}
	nums := parseInput(string(file))

	parent := createNode(nums)
	fmt.Println(sumEntries(parent, 0))
	fmt.Println(sumEntries2(parent, 0))
}

func parseInput(input string) *stack.Stack {
	s := stack.New()
	numRe := regexp.MustCompile(`\d+`)
	numRaw := numRe.FindAllString(input, -1)
	for i := len(numRaw) - 1; i >= 0; i-- {
		num, _ := strconv.Atoi(numRaw[i])
		s.Push(num)
	}

	return s
}

func createNode(in *stack.Stack) *node {
	n := node{childLen: in.Pop().(int), entryLen: in.Pop().(int)}
	for i := 0; i < n.childLen; i++ {
		n.Childs = append(n.Childs, createNode(in))
	}
	for i := 0; i < n.entryLen; i++ {
		n.Entries = append(n.Entries, in.Pop().(int))
	}

	return &n
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func sumEntries(p *node, s int) int {
	if p.childLen < 1 {
		return s + sum(p.Entries)
	}
	for i := 0; i < p.childLen; i++ {
		s = sumEntries(p.Childs[i], s)
	}
	return s + sum(p.Entries)
}

func sumEntries2(p *node, s int) int {
	if p.childLen < 1 {
		return s + sum(p.Entries)
	}
	for i := 0; i < p.entryLen; i++ {
		if p.Entries[i] <= p.childLen {
			s = sumEntries2(p.Childs[p.Entries[i]-1], s)
		}
	}
	return s
}
