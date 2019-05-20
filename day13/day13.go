package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

const (
	filename = "input.txt"

	horizontal = 45  // -
	vertical   = 124 // |
	curveL     = 92  // \
	curveR     = 47  // /
	cross      = 43  // +

	cartDisp = "^>v<"
)

type direction int

const (
	up direction = iota
	right
	down
	left
)

type cart struct {
	X        int
	Y        int
	Dir      direction
	CrossDir int
	Crashed  bool
}

func newCart(x, y int, dir direction) *cart {
	return &cart{X: x, Y: y, Dir: dir}
}

func main() {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("file not found")
	}
	tracks, carts := parseInput(file)

	for len(carts) > 1 {
		sort.Slice(carts, func(i, k int) bool {
			if carts[i].Y != carts[k].Y {
				return carts[i].Y < carts[k].Y
			}
			return carts[i].X < carts[k].X
		})

		for i := range carts {
			switch carts[i].Dir {
			case up:
				carts[i].Y--
			case down:
				carts[i].Y++
			case left:
				carts[i].X--
			case right:
				carts[i].X++
			}

			switch tracks[carts[i].Y][carts[i].X] {
			case curveL:
				switch carts[i].Dir {
				case up:
					carts[i].Dir = left
				case down:
					carts[i].Dir = right
				case left:
					carts[i].Dir = up
				case right:
					carts[i].Dir = down
				}
			case curveR:
				switch carts[i].Dir {
				case up:
					carts[i].Dir = right
				case down:
					carts[i].Dir = left
				case left:
					carts[i].Dir = down
				case right:
					carts[i].Dir = up
				}
			case cross:
				carts[i].Dir = mod(carts[i].Dir+direction(carts[i].CrossDir-1), 4)
				carts[i].CrossDir = (carts[i].CrossDir + 1) % 3
			}

			for ic, c := range carts {
				if carts[i].X == c.X && carts[i].Y == c.Y {
					if ic == i {
						continue
					}
					fmt.Printf("Crash: %d,%d\n", c.X, c.Y)
					carts[i].Crashed = true
					carts[ic].Crashed = true
					break
				}
			}
		}
		for i := len(carts) - 1; i >= 0; i-- {
			if carts[i].Crashed {
				carts = append(carts[:i], carts[i+1:]...)
			}
		}
	}
	fmt.Printf("%d,%d\n", carts[0].X, carts[0].Y)
}

func parseInput(input []byte) ([][]byte, []cart) {
	var tracks [][]byte
	var carts []cart
	for y, line := range bytes.Split(input, []byte("\n")) {
		tracks = append(tracks, make([]byte, len(line)))
		for x := range line {
			isCard, dir := isCart(line[x])
			if isCard {
				carts = append(carts, *newCart(x, y, dir))
				if dir == up || dir == down {
					tracks[y][x] = '|'
				} else {
					tracks[y][x] = '-'
				}
			} else {
				tracks[y][x] = line[x]
			}
		}
	}

	return tracks, carts
}

func mod(d direction, m int) direction {
	res := d % direction(m)
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + direction(m)
	}
	return res
}

func isCart(b byte) (bool, direction) {
	for i := range cartDisp {
		if b == cartDisp[i] {
			return true, direction(i)
		}
	}
	return false, -1
}

func printTracks(tracks [][]byte, carts []cart) {
	for y := range tracks {
		for x := range tracks[y] {
			isCard, i := cartAtPos(x, y, carts)
			if isCard {
				fmt.Printf("%c", cartDisp[carts[i].Dir])
			} else {
				fmt.Printf("%c", tracks[y][x])
			}
		}
		fmt.Println()
	}
}

func cartAtPos(x, y int, carts []cart) (bool, int) {
	for i, cart := range carts {
		if cart.X == x && cart.Y == y {
			return true, i
		}
	}
	return false, -1
}
