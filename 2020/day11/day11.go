package main

import (
	"fmt"
	"strings"
)

const OCCUPIED = '#'
const FLOOR = '.'
const EMPTY = 'L'


func seatingRules(s []string) []string {
	chart := make([]string, len(s))

	for row, line := range s {
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
				}else{
				if i == 0 && line[i+1] != OCCUPIED || i == len(line)-1 && line[i-1] != OCCUPIED {
					newLine += "#"
				} else {
					newLine += string(line[i])
				}
				}

			case OCCUPIED:
				if adjCount(line, i){
					newLine += "L"
				}
				newLine += string(line[i])

			case FLOOR:
				newLine += "."
			}
		}
		chart[row] = newLine
	}
	return chart
}
func inbounds(s string, index int) bool {
	if index != 0 && index != len(s)-1 {
		return true
	}
	return false
}
func adjCount(s string, index int)bool{
	rightcount := 0
	leftcount := 0
	ptr := index -1

 	for ptr !=0{
		if s[ptr] == EMPTY{
			leftcount++
			ptr--
		}
	}
	for index != len(s)-1{
		if s[index] == EMPTY{
			rightcount++
			index++
		}
	}
	
	if rightcount >= 4 || leftcount >= 4 {
		return true
	}
	
	return false
}

func findState(m map[string]bool, seats []string) string{
	occupiedCount := 0
	chart := ""

	for _, row := range seats{
		chart += row
	}
	if m[chart]{
		for _, seat := range chart{
			if seat == OCCUPIED{
				occupiedCount++
			}
		return fmt.Sprintf("Number of Occupied Seats are %d", occupiedCount)
		}
	}
		m[chart] = true
		chart = ""
		newseats := seatingRules(seats)
		return findState(m, newseats)

}
func main() {
	input := strings.Split(data, "\n")
	ans := seatingRules(input)
	for _, r := range ans{
	fmt.Println(len(r))
	}
fmt.Println(ans)
}
