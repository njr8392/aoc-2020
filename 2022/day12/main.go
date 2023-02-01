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

func main() {
	input := util.ReadInput("input.txt")
	grid := bytes.Split(input, []byte("\n"))
	fmt.Println(path(grid))
}

func path(grid [][]byte) int {
	R := len(grid)    // Row bounds
	C := len(grid[0]) // Column bounds
	var q []int
	seen := make(map[[2]int]bool)
	step := 0

	//find starting posistion
	r, c := FindStart(grid)
	grid[r][c] = byte('a')

	q = append(q, r, c, 0)
	seen[[2]int{r, c}] = true

	for len(q) > 0 {
		r = q[0]
		c = q[1]
		step = q[2]
		fmt.Println(r, c, string(grid[r][c]))
		q = q[3:]

		if grid[r][c] == byte('E') {
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

			if grid[next_r][next_c] > grid[r][c]+1 {
				continue
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
