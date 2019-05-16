package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const filename = "input.txt"

type pair struct {
	X, Y int
}

type point struct {
	Pos, Vel pair
}

func main() {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("file not found")
	}
	points := parseInput(string(file))

	minByte := 1<<63 - 1

	for sec := 0; sec < 100000; sec++ {
		minPos := pair{X: 1<<63 - 1, Y: 1<<63 - 1}
		maxPos := pair{X: -1 << 63, Y: -1 << 63}

		for _, point := range points {
			if point.Pos.X > maxPos.X {
				maxPos.X = point.Pos.X
			}
			if point.Pos.Y > maxPos.Y {
				maxPos.Y = point.Pos.Y
			}
			if point.Pos.X < minPos.X {
				minPos.X = point.Pos.X
			}
			if point.Pos.Y < minPos.Y {
				minPos.Y = point.Pos.Y
			}
		}

		size := (maxPos.Y - minPos.Y) * (maxPos.X - minPos.X)
		if size < minByte {
			minByte = size
		} else {

			// go back one step
			for i := range points {
				points[i].Pos.X -= points[i].Vel.X
				points[i].Pos.Y -= points[i].Vel.Y
			}

			minPos := pair{X: 1<<63 - 1, Y: 1<<63 - 1}
			maxPos := pair{X: -1 << 63, Y: -1 << 63}

			for _, point := range points {
				if point.Pos.X > maxPos.X {
					maxPos.X = point.Pos.X
				}
				if point.Pos.Y > maxPos.Y {
					maxPos.Y = point.Pos.Y
				}
				if point.Pos.X < minPos.X {
					minPos.X = point.Pos.X
				}
				if point.Pos.Y < minPos.Y {
					minPos.Y = point.Pos.Y
				}
			}

			field := make([][]byte, maxPos.Y-minPos.Y+1)
			for i := range field {
				field[i] = make([]byte, maxPos.X-minPos.X+1)
				for k := range field[i] {
					field[i][k] = '.'
				}
			}

			for _, point := range points {
				field[point.Pos.Y-minPos.Y][point.Pos.X-minPos.X] = '#'
			}

			for i := range field {
				for k := range field[i] {
					fmt.Printf("%c", field[i][k])
				}
				fmt.Println()
			}
			fmt.Println()

			fmt.Println(sec - 1)
			break
		}

		for i := range points {
			points[i].Pos.X += points[i].Vel.X
			points[i].Pos.Y += points[i].Vel.Y
		}
	}
}

func parseInput(input string) []point {
	var points []point
	numRe := regexp.MustCompile(`(-)*\d+`)

	for _, line := range strings.Split(input, "\n") {
		numsRaw := numRe.FindAllString(line, 4)
		var nums []int
		for _, num := range numsRaw {
			n, _ := strconv.Atoi(num)
			nums = append(nums, n)
		}
		p := point{Pos: pair{X: nums[0], Y: nums[1]}, Vel: pair{X: nums[2], Y: nums[3]}}
		points = append(points, p)
	}

	return points
}
