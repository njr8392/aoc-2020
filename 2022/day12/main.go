package main

import (
	"bytes"
	"fmt"
	"github.com/njr8392/aoc/util"
)

var (
	dr []int = []int{-1, 1, 0, 0}
	dc []int = []int{0, 0, 1, -1}
)

const (
	START = 'a'
)

func main() {
	input := util.ReadInput("input.txt")
	grid := bytes.Split(input, []byte("\n"))
	r, c := FindStart(grid)
	fmt.Println(path(grid, r, c, 'E'))	//part 1
	r, c = FindEnd(grid)
	fmt.Println(path(grid, r, c, 'a'))	//part 2
}

/* path will take the starting posistion and the byte that triggers the search to stop.
It adjust the starting posistion of the grid depending on the part of the problem to solve.
For part 1, will we set the starting posistion to the posistion of 'S' and walking until 
we arrive at 'E'. Adjusting for part 2, the start and end positions are flipped.
Once the posistion is set, path will walk by BFS and count the number of steps */
func path(grid [][]byte, r, c int, end byte) int {
	R := len(grid)    // Row bounds
	C := len(grid[0]) // Column bounds
	var q []int
	seen := make(map[[2]int]bool)
	step := 0

	//set the starting byte for the search
	grid[r][c] = byte('a')
	if end == START {
		grid[r][c] = byte('z')
	}

	q = append(q, r, c, 0)
	seen[[2]int{r, c}] = true

	for len(q) > 0 {
		r = q[0]
		c = q[1]
		step = q[2]
		q = q[3:]

		if grid[r][c] == end {
			return step
		}

		//get neighbors and check bounds
		for i := 0; i < 4; i++ {
			next_r := r + dr[i]
			next_c := c + dc[i]

			if next_r >= R || next_r < 0 {
				continue
			}
			if next_c >= C || next_c < 0 {
				continue
			}

			//cannot walk where the next byte is greater than 1 from the previous byte
			//checks are conditional dependent on the direction we are walking
			if end == 'E' {
				if grid[next_r][next_c] > grid[r][c]+1 {
					continue
				}
			} else {
				if grid[next_r][next_c]+1 < grid[r][c] {
					continue
				}
			}

			if !seen[[2]int{next_r, next_c}] {
				seen[[2]int{next_r, next_c}] = true
				q = append(q, next_r, next_c, step+1)
			}

		}

	}
	return 0
}

func FindStart(b [][]byte) (int, int) {
	for i := range b {
		for j := range b[i] {
			if b[i][j] == byte('S') {
				return i, j
			}
		}
	}
	return 0, 0
}

func FindEnd(b [][]byte) (int, int) {
	for i := range b {
		for j := range b[i] {
			if b[i][j] == byte('E') {
				return i, j
			}
		}
	}
	return 0, 0
}
