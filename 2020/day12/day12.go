package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type ship struct {
	waypoint, distance            point
	posistion                     rune
	north, south, east, west, deg int
}

type point struct {
	x, y int
}

func (s *ship) Instruction(instr string) {
	dir := instr[0]
	val := instr[1:]

	num, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println(err)
	}

	switch dir {
	case 'N':
		s.north += num

	case 'S':
		s.south += num

	case 'E':
		s.east += num
	case 'W':
		s.west += num
	case 'F':
		s.Instruction(string(s.posistion) + val)

	default:
		s.Rotate(rune(dir), num)
	}
}

func (s *ship) Rotate(direction rune, degree int) {
	switch direction {
	case 'R':
		s.deg += degree

	case 'L':
		s.deg -= degree
	}

	s.Abs()
	s.ChangeDirection()
}

func (s *ship) ChangeDirection() {
	switch s.deg {
	case 0:
		s.posistion = 'N'
	case 90:
		s.posistion = 'E'
	case 180:
		s.posistion = 'S'
	case 270:
		s.posistion = 'W'
	}
}
func (s *ship) Abs() {
	if s.deg >= 360 {
		s.deg -= 360
	} else if s.deg < 0 {
		s.deg += 360
	} else {
		s.deg = s.deg
	}
}
func (s *ship) ManhattanSum() int {
	n := abs(s.north - s.south)
	e := abs(s.east - s.west)

	return n + e
}

func (s *ship) RotateCordinates(degree int) {
	rad := float64(degree) * math.Pi / float64(180)
	tx := float64(s.waypoint.x)*math.Cos(float64(rad)) - float64(s.waypoint.y)*math.Sin(float64(rad))
	ty := float64(s.waypoint.y)*math.Cos(float64(rad)) + float64(s.waypoint.x)*math.Sin(float64(rad))
	tx = math.Round(tx)
	ty = math.Round(ty)
	s.waypoint = point{int(tx), int(ty)}

}

func (s *ship) Part2(instr string) {
	dir := instr[0]
	val := instr[1:]

	num, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println(err)
	}

	switch dir {
	case 'N':
		s.waypoint.y += num

	case 'S':
		s.waypoint.y -= num

	case 'E':
		s.waypoint.x += num
	case 'W':
		s.waypoint.x -= num
	case 'F':
		s.distance.x += s.waypoint.x * num
		s.distance.y += s.waypoint.y * num
	case 'L':
		s.RotateCordinates(num)
	case 'R':
		s.RotateCordinates(-num)
	}
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func main() {
	sh := ship{posistion: 'E', deg: 90}
	input := strings.Split(data, "\n")
	for _, pos := range input {
		sh.Instruction(pos)

	}
	//part 1
	fmt.Println(sh.ManhattanSum())

	sh2 := ship{waypoint: point{10, 1}}
	for i, pos := range input {
		sh2.Part2(pos)
		if i < 15 {

			fmt.Printf("Waypoint %v\t distance: %v\n", sh2.waypoint, sh2.distance)
		}

	}
	//part 2
	fmt.Println(abs(sh2.distance.x) + abs(sh2.distance.y))
}
