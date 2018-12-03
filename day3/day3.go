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

type claim struct {
	ID     int
	Left   int
	Top    int
	Width  int
	Height int
}

func main() {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("file not found")
	}
	lines := strings.Split(string(file), "\n")

	// part1(lines)
	part2(lines)
}

func newClaim(data string) claim {
	nums := make([]int, 5)

	re := regexp.MustCompile(`\d+`)
	res := re.FindAllString(data, -1)

	for i := range res {
		nums[i], _ = strconv.Atoi(res[i])
	}

	return claim{
		ID:     nums[0],
		Left:   nums[1],
		Top:    nums[2],
		Width:  nums[3],
		Height: nums[4]}
}

func part1(data []string) {
	var fabric [1000][1000]int
	var overlap int

	for _, line := range data {
		c := newClaim(line)
		for x := c.Left; x < c.Left+c.Width; x++ {
			for y := c.Top; y < c.Top+c.Height; y++ {
				fabric[x][y]++
			}
		}
	}

	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if fabric[x][y] > 1 {
				overlap++
			}
		}
	}
	fmt.Println(overlap)
}

func part2(data []string) {
	var fabric [1000][1000]int
	var claims []claim

	for _, line := range data {
		c := newClaim(line)
		claims = append(claims, c)

		for x := c.Left; x < c.Left+c.Width; x++ {
			for y := c.Top; y < c.Top+c.Height; y++ {
				fabric[x][y]++
			}
		}
	}
	for _, c := range claims {
		if c.checkClaim(&fabric) {
			fmt.Println(c.ID)
		}
	}
}

func (c *claim) checkClaim(fabric *[1000][1000]int) bool {
	for x := c.Left; x < c.Left+c.Width; x++ {
		for y := c.Top; y < c.Top+c.Height; y++ {
			if fabric[x][y] != 1 {
				return false
			}
		}
	}
	return true
}
