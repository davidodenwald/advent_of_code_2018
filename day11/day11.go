package main

import "fmt"

const serial = 7672

func main() {
	grid := [300][300]int{}
	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			rackID := x + 10
			powerLvl := (rackID*y + serial) * rackID
			powerLvl = (powerLvl / 100) % 10
			powerLvl -= 5
			grid[x-1][y-1] = powerLvl
		}
	}
	max := -1 << 63
	var maxX, maxY, maxS int
	for s := 1; s <= 300; s++ {
		for x := 1; x <= 300; x++ {
			if x > 300-s+1 {
				continue
			}
			for y := 1; y <= 300; y++ {
				if y > 300-s+1 {
					continue
				}
				sum := sumGrid(&grid, x-1, y-1, s)
				if sum > max {
					max = sum
					maxX = x
					maxY = y
					maxS = s
				}
			}
		}
	}
	fmt.Printf("<%d,%d,%d>\n", maxX, maxY, maxS)
}

func sum3x3(grid *[300][300]int, x, y int) int {
	sum := 0
	for tx := 0; tx < 3; tx++ {
		for ty := 0; ty < 3; ty++ {
			sum += grid[x+tx][y+ty]
		}
	}
	return sum
}

func sumGrid(grid *[300][300]int, x, y, s int) int {
	sum := 0
	for tx := 0; tx < s; tx++ {
		for ty := 0; ty < s; ty++ {
			sum += grid[x+tx][y+ty]
		}
	}
	return sum
}
