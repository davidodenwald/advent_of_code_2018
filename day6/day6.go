package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const filename = "input.txt"
const desiredDist = 10000

type point struct {
	id int
	x  int
	y  int
}

func main() {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("file not found")
	}

	var points []point
	var xMax int
	var yMax int
	for i, line := range strings.Split(string(file), "\n") {
		xy := strings.Split(line, ", ")
		x, _ := strconv.Atoi(xy[1])
		y, _ := strconv.Atoi(xy[0])
		points = append(points, point{
			id: i + 1,
			x:  x,
			y:  y})
		if x > xMax {
			xMax = x
		}
		if y > yMax {
			yMax = y
		}
	}
	grid := make([][]int, xMax+1)
	for i := range grid {
		grid[i] = make([]int, yMax+1)
	}

	// part2 doesn't work if part1 runs before
	// part1(points, grid)
	part2(points, grid)
}

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}

func dist(ax, ay, bx, by int) int {
	return abs(ax-bx) + abs(ay-by)
}

func part1(points []point, grid [][]int) {
	for _, p := range points {
		grid[p.x][p.y] = p.id
	}

	for x := range grid {
		for y := range grid[x] {

			minDist := 1<<63 - 1
			var minPoint int
			for _, p := range points {
				if p.x == x && p.y == y {
					minPoint = p.id
					break
				}
				currDist := dist(p.x, p.y, x, y)
				if currDist < minDist {
					minDist = currDist
					minPoint = p.id
				} else if currDist == minDist {
					minPoint = 0
				}
			}
			grid[x][y] = minPoint
		}
	}

	for x := range grid {
		for y := range grid[x] {
			if ((y == 0 || y == len(grid[x])-1) || (x == 0 || x == len(grid)-1)) && grid[x][y] != 0 {
				for i, p := range points {
					if p.id == grid[x][y] {
						points = append(points[:i], points[i+1:]...)
					}
				}
			}
		}
	}
	areas := make(map[int]int)
	for _, p := range points {
		for x := range grid {
			for y := range grid[x] {
				if grid[x][y] != p.id {
					continue
				}
				areas[p.id]++
			}
		}
	}
	var maxArea int
	for _, area := range areas {
		if area > maxArea {
			maxArea = area
		}
	}
	fmt.Println(maxArea)
}

func part2(points []point, grid [][]int) {
	var area int
	for x := range grid {
		for y := range grid[x] {
			var totalDist int
			for _, p := range points {
				totalDist += dist(p.x, p.y, x, y)
			}
			if totalDist < desiredDist {
				area++
			}
		}
	}
	fmt.Println(area)
}
