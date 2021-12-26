package main

import (
	"bytes"
	"fmt"
	"github.com/njr8392/aoc/util"
)

func main() {
	data := util.ReadInput("input.txt")
	grid := MakeGrid(data)
	path := LowRiskPath(grid)
	fmt.Printf("Answer to part 1 is %d\n", path)
}

//part 1
func LowRiskPath(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				if grid[i-1][j] > grid[i][j-1] {
					grid[i][j] += grid[i][j-1]
				} else {
					grid[i][j] += grid[i-1][j]
				}
			}

		}
	}
	return grid[len(grid)-1][len(grid[0])-1] - grid[0][0]
}

func MakeGrid(data []byte) [][]int {
	lines := bytes.Split(data, []byte{'\n'})
	lines = lines[:len(lines)-1] //idk last row is empty so drop it
	grid := make([][]int, len(lines))

	for i, row := range lines {
		grow := make([]int, len(row))
		for j, item := range row {
			grow[j] = ByteToInt(item)
		}
		grid[i] = grow
	}
	return grid
}

func ByteToInt(b byte) int {
	return int(b - '0')
}
