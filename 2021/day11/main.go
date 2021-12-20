package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type point struct {
	x, y int
}

func main() {
	data := ReadInput("input.txt")
	grid := ByteToInts(data)
	count := Sim(grid, 100)
	fmt.Printf("Answer to Part 1 is %d\n", count)
}
func Sim(grid [][]int, length int) int {
	var flash int
	for length > 0 {
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				grid[i][j]++
			}
		}
		count := 0
		grid, count = Flash(grid)
		flash += count
		length--
	}
	return flash
}
func Flash(grid [][]int) ([][]int, int) {
	seen := make(map[point]bool)
	flashcount := 0

	for {
		//flashcount counts the numbers of "flashes" and count's any instance of nums >9. if no, exist then we exit the loop
		count := 0
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				if grid[i][j] > 9 {
					p := point{i, j}

					if !seen[p] {
						grid = update(grid, i, j)
						flashcount++
						count++
					}
					seen[p] = true
				}
			}
		}
		if count == 0 {
			break
		}
	}
	//set all flash points to 0
	for key, val := range seen {
		if val {
			grid[key.x][key.y] = 0
		}
	}
	return grid, flashcount
}
func update(grid [][]int, i, j int) [][]int {
	var left, right, up, down int
	//set the bounds. only want to change the points directly surronding the current point in the grid
	if i-1 > 0 {
		up = i - 1
	}

	if i+1 > len(grid)-1 {
		down = len(grid) - 1
	} else {
		down = i + 1
	}

	if j-1 > 0 {
		left = j - 1
	}

	if j+1 > len(grid[0])-1 {
		right = len(grid[0]) - 1
	} else {
		right = j + 1
	}

	var walk func([][]int, int, int) bool
	seen := make(map[point]bool)
	//walk points around it and update on flash
	walk = func(grid [][]int, i, j int) bool {
		if i < up || i > down || j < left || j > right {
			return false
		}
		point := point{i, j}
		if seen[point] {
			return false
		}

		seen[point] = true
		grid[i][j] += 1

		return walk(grid, i-1, j) || walk(grid, i+1, j) || walk(grid, i, j-1) || walk(grid, i, j+1)
	}

	walk(grid, i, j)
	return grid
}
func ByteToInts(data []byte) [][]int {
	var grid [][]int
	buf := bytes.NewBuffer(data)
	for {
		line, err := buf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		row := make([]int, len(line)-1)
		for i, b := range line[:len(line)-1] {
			row[i] = int(b - '0')
		}
		grid = append(grid, row)
	}
	return grid
}

func ReadInput(file string) []byte {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, info.Size())
	_, err = f.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	return buf
}
