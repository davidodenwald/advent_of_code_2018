package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const filename = "input.txt"

type logLine struct {
	time  time.Time
	ID    int
	sleep bool
	wake  bool
}

func main() {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("file not found")
	}
	lines := strings.Split(string(file), "\n")

	var logs []logLine
	dateRe := regexp.MustCompile(`\d+-\d+-\d+ \d+:\d+`)
	IDRe := regexp.MustCompile(`#\d+`)
	for _, line := range lines {
		dateStr := dateRe.FindString(line)
		date, err := time.Parse("2006-01-02 15:04", dateStr)
		if err != nil {
			log.Printf("Could not parse date: %s\n", dateStr)
			continue
		}
		idStr := IDRe.FindString(line)
		var id int
		if len(idStr) > 0 {
			id, err = strconv.Atoi(idStr[1:])
			if err != nil {
				log.Printf("Could not parse guard id: %s\n", line)
				continue
			}
		}

		logs = append(logs, logLine{
			time:  date,
			ID:    id,
			sleep: strings.Contains(line, "falls asleep"),
			wake:  strings.Contains(line, "wakes up")})
	}
	sort.Slice(logs, func(i, j int) bool {
		return logs[i].time.Sub(logs[j].time) < 0
	})
	part1(logs)
	part2(logs)
}

func part1(logs []logLine) {
	guards := make(map[int]int) // [ID]sleep-min

	var currGuard int
	var sleepStart time.Time
	var max int
	var maxGuard int
	for i, line := range logs {
		if line.ID > 0 {
			currGuard = line.ID
		}
		if line.sleep {
			logs[i].ID = currGuard
			sleepStart = line.time
		}
		if line.wake {
			logs[i].ID = currGuard
			guards[currGuard] += int(line.time.Sub(sleepStart).Minutes())
			if guards[currGuard] > max {
				max = guards[currGuard]
				maxGuard = currGuard
			}
		}
	}

	mins := make(map[int]int) // [minute]sleep
	var minStart int
	max = 0
	var maxMin int
	for _, line := range logs {
		if line.ID == maxGuard && line.sleep {
			minStart = line.time.Minute()
		}
		if line.ID == maxGuard && line.wake {
			for i := minStart; i < line.time.Minute(); i++ {
				mins[i]++
				if mins[i] > max {
					max = mins[i]
					maxMin = i
				}
			}
		}
	}
	fmt.Println(maxGuard * maxMin)
}

func part2(logs []logLine) {
	var currGuard int
	var guards []int
	for i, l := range logs {
		if l.ID > 0 {
			currGuard = l.ID
			guards = append(guards, currGuard)
		} else {
			logs[i].ID = currGuard
		}
	}

	guardMaxTime := make(map[int]int) // [id]max-sleep
	var maxTime int
	var maxMin int
	var maxGuard int
	for _, guard := range guards {
		mins := make(map[int]int) // [minute]sleep
		var minStart int
		var max int
		for _, l := range logs {
			if l.ID == guard && l.sleep {
				minStart = l.time.Minute()
			}
			if l.ID == guard && l.wake {
				for i := minStart; i < l.time.Minute(); i++ {
					mins[i]++
					if mins[i] > max {
						max = mins[i]
						guardMaxTime[guard] = max
						maxMin = i
					}
				}
			}
		}
		if guardMaxTime[guard] > maxTime {
			maxTime = guardMaxTime[guard]
			maxGuard = guard
		}
	}
	fmt.Println(maxGuard * maxMin)
}
