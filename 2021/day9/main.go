package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

const (
	LOW    = 0
	UPPER  = 100
	MAXINT = 1<<63 - 1
)

func main() {
	data := ReadInput("input.txt")
	hmap := ConfigHeatMap(data)
	lowpoints, pos := FindLowPoints(hmap)
	p1 := sum(lowpoints) + len(lowpoints)
	fmt.Printf("Answer to part 1: %d\n", p1)

	var counter []int
	for _, p := range pos {
		size := WalkLowPoints(hmap, p)
		counter = append(counter, size)
	}
	sort.Ints(counter)
	size := len(counter) - 1
	fmt.Printf("Answer to part 2: %d\n", counter[size]*counter[size-1]*counter[size-2])
}

//coardinate in the grid of the heatmap
type point struct {
	x int
	y int
}

//walks the area around the lowpoints and counts the number of vaild calls ie. the size of the "basin"
func WalkLowPoints(grid [][]int, pos point) int {
	var walk func([][]int, int, int) bool
	var count int
	seen := make(map[point]bool)

	walk = func(grid [][]int, i, j int) bool {
		if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[0])-1 {
			return false
		}

		if grid[i][j] == 9 {
			return false
		}

		if seen[point{i, j}] {
			return false
		}
		seen[point{i, j}] = true
		count++
		return walk(grid, i-1, j) || walk(grid, i+1, j) || walk(grid, i, j-1) || walk(grid, i, j+1)
	}
	walk(grid, pos.x, pos.y)
	return count
}

//part 1. returns list of points for part 2 as well
func FindLowPoints(data [][]int) ([]int, []point) {
	var lowpoints []int
	var points []point
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {

			up := Up(data, i, j)
			down := Down(data, i, j)
			left := Left(data, i, j)
			right := Right(data, i, j)
			min := Min(up, down, left, right)

			if data[i][j] < min {
				point := point{i, j}
				lowpoints = append(lowpoints, data[i][j])
				points = append(points, point)
			}
		}
	}
	return lowpoints, points
}

//checks if the int directly above the current posistion is in the grid and return it.
//Otherwise we will return the max int (since we are only trying to find the lowest int with each valid adjacent move
//same logic will be applied to Down, Left, and Right functions
func Up(data [][]int, i, j int) int {
	if i-1 >= LOW {
		return data[i-1][j]
	}
	return MAXINT
}

func Down(data [][]int, i, j int) int {
	if i+1 <= len(data)-1 {
		return data[i+1][j]
	}
	return MAXINT
}

func Left(data [][]int, i, j int) int {
	if j-1 >= LOW {
		return data[i][j-1]
	}
	return MAXINT
}

func Right(data [][]int, i, j int) int {
	if j+1 <= len(data[0])-1 {
		return data[i][j+1]
	}
	return MAXINT
}

func ConfigHeatMap(data []byte) [][]int {
	buf := bytes.NewBuffer(data)
	var heatmap [][]int
	for {
		line, err := buf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		ints := make([]int, len(line)-1)
		for i, b := range line[:len(line)-1] {
			ints[i] = ByteToInt(b)
		}
		heatmap = append(heatmap, ints)
	}
	return heatmap
}

func ReadInput(file string) []byte {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	buf := make([]byte, info.Size())
	_, err = f.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
		return nil
	}
	return buf
}

func Min(nums ...int) int {
	min := MAXINT
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func ByteToInt(b byte) int {
	return int(b - '0')
}

func sum(nums []int) int {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	return sum
}
