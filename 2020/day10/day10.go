package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func readData(s string) []int {
	input := strings.Split(s, "\n")
	jolts := make([]int, len(input))

	for i, joltString := range input {
		jolt, err := strconv.Atoi(joltString)

		if err != nil {
			fmt.Printf("Error converting string to jolt(num): %s", err)
		}

		jolts[i] = jolt
	}
	return jolts
}

func joltDiffCount(j []int) (int, int) {
	j = append(j, 0)
	sorted := j[:]
	sort.Ints(sorted)
	diffcount := make(map[int]int)

	for i, jolt := range sorted {
		if i != 0 && i != len(j)-1 {
			diff := jolt - sorted[i-1]
			diffcount[diff]++
		}
		if i == 0 {
			diff := sorted[i+1] - 0
			diffcount[diff]++
		}

		if i == len(j)-1 {
			diffcount[3]++
		}
	}

	part2 := numConnections(sorted)
	return diffcount[1] * diffcount[3], part2
}

func numConnections(j []int) int {
	paths := make([]int, len(j))
	paths[0]++

	for i := 0; i < len(j); i++ {
		for k := i + 1; k < len(j); k++ {
			diff := j[k] - j[i]
			if diff > 3 {
				break
			}
			paths[k] += paths[i]
		}
	}

	return paths[len(paths)-1]
}

func main() {
	input := readData(data)
	part1, p2 := joltDiffCount(input)
	fmt.Printf("Solution to part 1 is %d\nSolution to Part 2 is %d\n", part1, p2)
}
