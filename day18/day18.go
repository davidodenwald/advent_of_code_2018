package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

const (
	filename = "input.txt"
	p1Rounds = 10
	rounds   = 1000000000

	ground = '.'
	tree   = '|'
	lumber = '#'
)

type pos struct {
	x, y int
}

func (p pos) adjacent() []pos {
	return []pos{
		{y: p.y - 1, x: p.x - 1},
		{y: p.y - 1, x: p.x},
		{y: p.y - 1, x: p.x + 1},
		{y: p.y, x: p.x - 1},
		{y: p.y, x: p.x + 1},
		{y: p.y + 1, x: p.x - 1},
		{y: p.y + 1, x: p.x},
		{y: p.y + 1, x: p.x + 1},
	}
}

func (p pos) inBounds(field [][]byte) bool {
	return p.y >= 0 && p.y < len(field) && p.x >= 0 && p.x < len(field[p.y])
}

func main() {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("file not found")
	}
	field := bytes.Split(file, []byte("\n"))

	stateMap := map[int][]int{}

	for i := 0; i < rounds; i++ {
		newField := *copyField(field)
		for y := range field {
			for x := range field[y] {
				newField[y][x] = getNext(pos{x: x, y: y}, field)
			}
		}
		field = newField

		trees := 0
		lumbers := 0
		for y := range field {
			for x := range field[y] {
				switch field[y][x] {
				case tree:
					trees++
				case lumber:
					lumbers++
				}
			}
		}

		if i == p1Rounds-1 {
			fmt.Println("Part 1:", trees*lumbers)
		}

		stateMap[trees*lumbers] = append(stateMap[trees*lumbers], i)
		if len(stateMap[trees*lumbers]) > 5 {
			break
		}
	}

	states := 0
	for _, v := range stateMap {
		if len(v) > 3 {
			states++
		}
	}
	for k, v := range stateMap {
		if len(v) > 3 {
			if v[0]%28 == (rounds-1)%28 {
				fmt.Println("Part 2:", k)
			}
		}
	}
}

func printField(field [][]byte) {
	for y := range field {
		for x := range field[y] {
			fmt.Printf("%c", field[y][x])
		}
		fmt.Println()
	}
	fmt.Println()
	time.Sleep(time.Millisecond * 66)
}

func copyField(field [][]byte) *[][]byte {
	cpy := make([][]byte, len(field))
	for y := range field {
		cpy[y] = make([]byte, len(field[y]))
		copy(cpy[y], field[y])
	}
	return &cpy
}

func getNext(p pos, field [][]byte) byte {
	trees := 0
	lumbers := 0
	for _, adj := range p.adjacent() {
		if !adj.inBounds(field) {
			continue
		}
		switch field[adj.y][adj.x] {
		case tree:
			trees++
		case lumber:
			lumbers++
		}
	}

	switch field[p.y][p.x] {
	case ground:
		if trees >= 3 {
			return tree
		}
		return ground

	case tree:
		if lumbers >= 3 {
			return lumber
		}
		return tree

	case lumber:
		if lumbers >= 1 && trees >= 1 {
			return lumber
		}
		return ground
	}
	return ' '
}
