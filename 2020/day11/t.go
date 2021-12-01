package main

import (
	"bytes"
	"fmt"
	"strings"
)

const OCCUPIED = '#'
const FLOOR = '.'
const EMPTY = 'L'

func seatingRules(s string) string {
	var chart bytes.Buffer
	ns := strings.Split(s, "\n")

	for _, line := range ns {
		newLine := ""
		for i, seat := range line {

			switch seat {
			case EMPTY:
				if inbounds(line, i) {
					if line[i-1] != OCCUPIED && line[i+1] != OCCUPIED {
						newLine += "#"
					} else {

						newLine += string(line[i])
					}
				} else {
					if i == 0 && line[i+1] != OCCUPIED || i == len(line)-1 && line[i-1] != OCCUPIED {
						newLine += "#"
					} else {
						newLine += string(line[i])
					}
				}

			case OCCUPIED:
				if adjCount(line, i) {
					newLine += "L"
				}
				newLine += string(line[i])

			case FLOOR:
				newLine += "."
			}
		}
		chart.WriteString(newLine + "\n")
	}
	return chart.String()
}
func inbounds(s string, index int) bool {
	if index != 0 && index != len(s)-1 {
		return true
	}
	return false
}
func adjCount(s string, index int) bool {
	rightcount := 0
	leftcount := 0
	ptr := index - 1

	for ptr >=  0 {
		if s[ptr] == EMPTY {
			leftcount++
			ptr--
		}
	}
	for index != len(s)-1 {
		if s[index] == EMPTY {
			rightcount++
			index++
		}
	}

	if rightcount >= 4 || leftcount >= 4 {
		return true
	}

	return false
}

func main() {
	var transform func(s string)
	seen := make(map[string]bool)

	transform = func(s string) {
		change := seatingRules(s)
		fmt.Println(change)

		if !seen[change] {
			seen[change] = true
			transform(change)
		}
		fmt.Println(change)
	}

	transform(data)
}
