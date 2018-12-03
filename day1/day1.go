package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const filename = "input.txt"

func main() {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("file not found")
	}

	var nums []int
	res := strings.Split(string(file), "\n")
	for x, line := range res {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Printf("%s:%d '%s' is not a number", filename, x+1, line)
		}
		nums = append(nums, num)
	}

	part1(nums)
	part2(nums)
}

func part1(nums []int) {
	var sum int
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
}

func part2(nums []int) {
	var sum int
	hist := make(map[int]bool)

	var i int
	for true {
		_, double := hist[sum]
		if double {
			fmt.Println(sum)
			return
		}
		hist[sum] = true
		sum += nums[i]
		i = (i + 1) % len(nums)
	}
}
