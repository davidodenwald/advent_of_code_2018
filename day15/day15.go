package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

const filename = "example.txt"

var field [][]byte

func main() {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("file not found")
	}
	field = bytes.Split(file, []byte("\n"))
	warriors := parseField()

	part1(warriors)
}

type position struct {
	X, Y int
}

func (p position) equal(pos position) bool {
	return p.X == pos.X && p.Y == pos.Y
}

func (p position) adjacent() []position {
	return []position{
		{p.X, p.Y - 1},
		{p.X - 1, p.Y},
		{p.X + 1, p.Y},
		{p.X, p.Y + 1},
	}
}

type warrior struct {
	Pos  position
	HP   int
	Elf  bool
	Dead bool
}

func newElf(p position) *warrior {
	return &warrior{Pos: p, HP: 200, Elf: true}
}

func newGoblin(p position) *warrior {
	return &warrior{Pos: p, HP: 200}
}

func (w *warrior) move(p position) {
	sym := field[w.Pos.Y][w.Pos.X]
	field[w.Pos.Y][w.Pos.X] = '.'
	w.Pos = p
	field[w.Pos.Y][w.Pos.X] = sym
}

func (w warrior) canAttack(enemies *[]warrior) (bool, *warrior) {
	var reach []*warrior
	var pos []position
	minHP := 1<<63 - 1
	for i, enemy := range *enemies {
		for _, p := range w.Pos.adjacent() {
			if w.Elf != enemy.Elf && p.equal(enemy.Pos) && !enemy.Dead {
				reach = append(reach, &(*enemies)[i])
				pos = append(pos, enemy.Pos)

				if enemy.HP < minHP {
					minHP = enemy.HP
				}
			}
		}
	}
	if len(reach) < 1 {
		return false, nil
	}

	for i := len(reach) - 1; i >= 0; i-- {
		if reach[i].HP > minHP {
			reach = append(reach[:i], reach[i+1:]...)
		}
	}

	sort.Slice(reach, func(i, k int) bool {
		if reach[i].Pos.Y != reach[k].Pos.Y {
			return reach[i].Pos.Y < reach[k].Pos.Y
		}
		return reach[i].Pos.X < reach[k].Pos.X
	})

	return true, reach[0]
}

func parseField() []warrior {
	var w []warrior

	for y := range field {
		for x := range field[y] {
			if field[y][x] == 'E' {
				w = append(w, *newElf(position{x, y}))
			}
			if field[y][x] == 'G' {
				w = append(w, *newGoblin(position{x, y}))
			}
		}
	}
	return w
}

func part1(warriors []warrior) {
	gameOver := false
	rounds := 0
	for !gameOver {
		sort.Slice(warriors, func(i, k int) bool {
			if warriors[i].Pos.Y != warriors[k].Pos.Y {
				return warriors[i].Pos.Y < warriors[k].Pos.Y
			}
			return warriors[i].Pos.X < warriors[k].Pos.X
		})
		for i := range warriors {
			if warriors[i].Dead {
				continue
			}
			attack, enemy := warriors[i].canAttack(&warriors)
			if !attack {
				validPos := inRange(warriors, warriors[i])
				validPos = reachable(warriors[i], validPos)
				if len(validPos) > 0 {
					stepPos := nextStep(warriors[i], validPos)
					warriors[i].move(stepPos)
				}
			}
			attack, enemy = warriors[i].canAttack(&warriors)
			if attack {
				enemy.HP -= 3
				if enemy.HP <= 0 {
					enemy.Dead = true
					field[enemy.Pos.Y][enemy.Pos.X] = '.'
				}
			}

			eAlive := 0
			gAlive := 0
			for i := len(warriors) - 1; i >= 0; i-- {
				if !warriors[i].Dead {
					if warriors[i].Elf {
						eAlive++
					} else {
						gAlive++
					}
				}
			}
			if eAlive < 1 {
				if i == len(warriors)-1 {
					rounds++
				}
				fmt.Println("Goblins win after round", rounds)
				gameOver = true
				break
			}
			if gAlive < 1 {
				if i == len(warriors)-1 {
					rounds++
				}
				fmt.Println("Elfs win after round", rounds)
				gameOver = true
				break
			}
		}

		for i := len(warriors) - 1; i >= 0; i-- {
			if warriors[i].Dead {
				warriors = append(warriors[:i], warriors[i+1:]...)
			}
		}

		fmt.Println("Round:", rounds)
		printField(warriors)
		if gameOver {
			hpSum := 0
			for _, w := range warriors {
				hpSum += w.HP
			}
			fmt.Println("HP left:", hpSum, "->", hpSum*rounds)
		}
		rounds++
	}
}

func printField(warriors []warrior) {
	for y := range field {
		for x := range field[y] {
			fmt.Printf("%c", field[y][x])
		}
		fmt.Print("\t")
		for _, w := range warriors {
			if w.Pos.Y == y {
				if w.Elf {
					fmt.Printf("E(%d) ", w.HP)
				} else {
					fmt.Printf("G(%d) ", w.HP)
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func inRange(warriors []warrior, w warrior) []position {
	var pos []position
	for _, war := range warriors {
		if w.Elf != war.Elf {
			for _, p := range war.Pos.adjacent() {
				if field[p.Y][p.X] == '.' {
					pos = append(pos, p)
				}
			}
		}
	}
	return pos
}

func reachable(w warrior, pos []position) []position {
	reach := map[position]bool{}
	sym := field[w.Pos.Y][w.Pos.X]
	field[w.Pos.Y][w.Pos.X] = '.'
	floodFill(w.Pos, &reach)
	field[w.Pos.Y][w.Pos.X] = sym

	var res []position
	for pReach := range reach {
		for _, pValid := range pos {
			if pReach.equal(pValid) {
				res = append(res, pValid)
			}
		}
	}
	return res
}

func floodFill(pos position, valid *map[position]bool) {
	if _, seen := (*valid)[pos]; seen {
		return
	}
	if field[pos.Y][pos.X] != '.' {
		return
	}
	(*valid)[pos] = true
	for _, p := range pos.adjacent() {
		floodFill(p, valid)
	}
}

func nextStep(start warrior, reachabl []position) position {
	dists := map[position]int{}
	for _, w := range reachabl {
		floodFillDist(w, &dists, 0)
	}

	min := 1<<63 - 1
	for pos, dist := range dists {
		for _, sPos := range start.Pos.adjacent() {
			if sPos.equal(pos) && dist < min {
				min = dist
			}
		}
	}

	var minPos []position
	for pos, dist := range dists {
		for _, sPos := range start.Pos.adjacent() {
			if sPos.equal(pos) && dist == min {
				minPos = append(minPos, pos)
			}
		}
	}

	sort.Slice(minPos, func(i, k int) bool {
		if minPos[i].Y != minPos[k].Y {
			return minPos[i].Y < minPos[k].Y
		}
		return minPos[i].X < minPos[k].X
	})

	return minPos[0]
}

func floodFillDist(pos position, valid *map[position]int, dist int) {
	if _, seen := (*valid)[pos]; seen && dist >= (*valid)[pos] {
		return
	}
	if field[pos.Y][pos.X] != '.' {
		return
	}
	(*valid)[pos] = dist
	for _, p := range pos.adjacent() {
		floodFillDist(p, valid, dist+1)
	}
}
