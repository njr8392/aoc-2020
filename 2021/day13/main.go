package main

import (
	"bytes"
	"fmt"
	"github.com/njr8392/aoc/util"
	"io"
	"log"
)

//fold along y --- len(matrix)/2
//fold along y --- len(matrix[0])/2
func main() {
	data := util.ReadInput("input.txt")
	chart := bytes.Split(data, []byte("\n\n"))

	// stupid hack
	chart[0] = append(chart[0], '\n')
	c, instruct := chart[0], chart[1]
	splits := ParseInstruction(instruct)

	set, err := SetPoints(c)
	if err != nil {
		fmt.Println(err)
	}
	grid := SetGrid(set) // to get part1 call fold once on grid
	
	//part2 
	grid = Execute(splits, grid)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				fmt.Print(".")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
func Execute(in []instruction, grid [][]int) [][]int {
	for _, i := range in {
		grid = Fold(grid, i.posistion, i.axis)
	}
	return grid
}

type point struct {
	x, y int
}

func Fold(grid [][]int, pos int, along string) [][]int {
	fmt.Println(len(grid), len(grid[0]), along, pos)
	if along == "y" {
		new_grid := Grid(len(grid[0]), len(grid)/2)
		//fold on y.  clone the top half as is
		for i := 0; i < pos; i++ {
			for j := 0; j < len(grid[0]); j++ {
				new_grid[i][j] |= grid[i][j]
			}
		}
		//work backwards  bottom left old = top left of new
		for i := len(grid) - 1; i > pos; i-- {
			for j := 0; j < len(grid[0]); j++ {
				new_grid[len(grid)-i-1][j] |= grid[i][j]
			}
		}
		return new_grid
	}
	new_grid := Grid(len(grid[0])/2, len(grid))
	//fold on y.  clone the top half as is
	for i := 0; i < len(grid); i++ {
		for j := 0; j < pos; j++ {
			new_grid[i][j] |= grid[i][j]
		}
	}
	//work backwards  top right old = top left of new
	for i := 0; i < len(grid); i++ {
		for j := len(grid[0]) - 1; j > pos; j-- {
			new_grid[i][len(grid[0])-j-1] |= grid[i][j]
		}
	}
	return new_grid

}

func SetPoints(data []byte) ([]point, error) {
	var set []point
	buf := bytes.NewBuffer(data)
	for {
		line, err := buf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		pos := bytes.Split(line[:len(line)-1], []byte{','})
		p := point{ByteToInt(pos[0]), ByteToInt(pos[1])}
		set = append(set, p)
	}
	return set, nil
}

type instruction struct {
	axis      string
	posistion int
}

func ParseInstruction(data []byte) []instruction {
	var list []instruction
	buf := bytes.NewBuffer(data)
	for {
		line, err := buf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		input := bytes.Split(line[:len(line)-1], []byte{' '})
		instruct := bytes.Split(input[2], []byte{'='})
		in := instruction{string(instruct[0]), ByteToInt(instruct[1])}
		list = append(list, in)
	}
	return list
}

func Grid(x, y int) [][]int {
	grid := make([][]int, y)
	for i := 0; i < len(grid); i++ {
		row := make([]int, x)
		grid[i] = row
	}
	return grid
}
func SetGrid(set []point) [][]int {
	xmax, ymax := Max(set)
	grid := Grid(xmax, ymax)
	for _, p := range set {
		grid[p.y][p.x] = 1
	}
	return grid
}

func ByteToInt(b []byte) int {
	var num int
	for _, x := range b {
		num = num*10 + int(x-'0')
	}
	return num
}

//add one to each max because the posistion count from 0. ie should be indexable
func Max(set []point) (xmax, ymax int) {
	for _, p := range set {
		if p.x+1 > xmax {
			xmax = p.x + 1
		}
		if p.y+1 > ymax {
			ymax = p.y + 1
		}
	}
	return xmax, ymax
}
