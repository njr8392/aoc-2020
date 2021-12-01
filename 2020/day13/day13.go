package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var earliest int = 1000000
	var bus int
	data := strings.Split(input, "\n")
	time, err := strconv.Atoi(data[0])
	if err != nil {
		fmt.Println(err)
	}
	buses := strings.Split(data[1], ",")
	for _, id := range buses {
		if id != "x" {
			b, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println(err)
			}
			diff := time % b
			nexttime := b - diff
			if nexttime < earliest {
				bus = b
				earliest = nexttime
			}
		}

	}

	part1 := earliest * bus

	p2 := part2(buses)
	fmt.Println(p2)
	fmt.Printf("solution to part 1: %d\nsolution to part2: %d\n", part1, p2)
}

func part2(buses []string) int {
	var schedule []int
	for _, id := range buses {
		if id != "x" {
			bus, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println(err)
			}
			schedule = append(schedule, bus)
		} else {
			schedule = append(schedule, 1)
		}
	}
	stamp := 1
	for {
		timeToSkipIfNoMatch := 1
		valid := true

		for i := 0; i < len(schedule); i++ {

			//check if bus id divides into the time
			if (stamp+i)%schedule[i] != 0 {
				valid = false
				break
			}

			//if so multiple by the next skip because it wont happen again for some time
			timeToSkipIfNoMatch *= schedule[i]
		}

		if valid {
			return stamp
		}

		stamp += timeToSkipIfNoMatch
	}
}
