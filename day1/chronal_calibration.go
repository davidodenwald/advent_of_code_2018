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

	sum := 0
	res := strings.Split(string(file), "\n")
	for x, line := range res {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("%s:%d '%s' is not a number", filename, x+1, line)
		}
		sum += num
	}
	fmt.Println(sum)
}
