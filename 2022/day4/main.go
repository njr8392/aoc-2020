package main

import (
	"bytes"
	"fmt"
	"github.com/njr8392/aoc/util"
	"strconv"
)

type point struct {
	x, y int
}

func main() {
	f := util.ReadInput("input.txt")
	part1 := 0
	part2 := 0
	lines := bytes.Split(f, []byte("\n"))
	for _, line := range lines {
		pair := bytes.Split(line, []byte(","))
		p1 := toPoint(pair[0])
		p2 := toPoint(pair[1])

		if check1(p1, p2) {
			part1++
		}
		if check2(p1, p2) {
			part2++
		}
	}
	fmt.Println(part1, part2)
}

func toPoint(b []byte) point {
	pt := bytes.Split(b, []byte("-"))
	fmt.Println(pt[0], pt[1])
	p1, _ := strconv.Atoi(string(pt[0]))
	p2, _ := strconv.Atoi(string(pt[1]))
	return point{x: p1, y: p2}
}

func check1(x, y point) bool {
	if (x.x <= y.x && x.y >= y.y) || (x.x >= y.x && x.y <= y.y) {
		return true
	}
	return false
}
func check2(x, y point) bool {
	if x.x > x.y {
		x.x, x.y, y.x, y.y = y.x, y.y, x.x, x.y
	}
	return x.y >= y.x
}
