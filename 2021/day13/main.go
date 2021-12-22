package main

import (
	"bytes"
	"fmt"
	"github.com/njr8392/aoc/util"
	"io"
)

//fold along y --- len(matrix)/2
//fold along y --- len(matrix[0])/2
func main() {
	data := util.ReadInput("input.txt")
	chart := bytes.Split(data, []byte("\n\n"))
	c, instruct := chart[0], chart[1]

	set, err := SetPoints(c)
	if err != nil {
		fmt.Println(err)
	}
	grid := SetGrid(set)
	fmt.Println(len(grid) - 447)
	fmt.Println(instruct)

	fold := Fold(grid, 655, "x")
//	f := Fold(fold, 447,"y")
//	fmt.Println(len(fold))
	// count dots
	count:=0
	for i:=0; i< len(fold);i++{
		for j:=0; j <len(fold[0]); j++{
			if fold[i][j] == 1{
			count++
			}
		}
	}
	fmt.Println(count)
}

type point struct {
	x, y int
}

func Fold(grid [][]int, pos int, along string) [][]int {
	if along == "y" {
		new_grid := Grid(len(grid[0]), len(grid)/2) 
		//fold on y.  clone the top half as is
		for i := 0; i < pos; i++ {
			for j := 0; j < len(grid[0]); j++ {
				new_grid[i][j] |= grid[i][j]
			}
		}
		fmt.Println(len(new_grid))
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
		fmt.Println(len("test"))
		//work backwards  bottom left old = top left of new
		for i := 0; i< len(grid); i++ {
			for j := len(grid[0])-1; j > pos; j-- {
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
