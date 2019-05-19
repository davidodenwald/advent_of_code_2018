package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

const (
	filename = "example.txt"

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
}

func newCart(x, y int, dir direction) *cart {
	return &cart{X: x, Y: y, Dir: dir, CrossDir: 0}
}

func main() {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("file not found")
	}
	tracks, carts := parseInput(file)

	crashed := false
	for !crashed {
		sort.Slice(carts, func(i, k int) bool {
			if carts[i].Y != carts[k].Y {
				return carts[i].Y < carts[k].Y
			}
			return carts[i].X < carts[k].X
		})
		for i := range carts {
			fmt.Printf("<%d,%d>\n", carts[i].X, carts[i].Y)
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
				carts[i].Dir = abs((carts[i].Dir + direction((carts[i].CrossDir%3)-1)) % 4)
				carts[i].CrossDir++
			}

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

			// printTracks(tracks, carts)

			for _, aCart := range carts {
				for _, bCart := range carts {
					if aCart == bCart {
						continue
					}
					if aCart.X == bCart.X && aCart.Y == bCart.Y {
						fmt.Printf("%d,%d\n", aCart.X, aCart.Y)
						crashed = true
					}
				}
			}
		}
		crashed = true
	}
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

func abs(n direction) direction {
	y := n >> 63
	return (n ^ y) - y
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
