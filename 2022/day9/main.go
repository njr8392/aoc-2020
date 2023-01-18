package main

import (
	"fmt"
	"github.com/njr8392/aoc/util"
	"bytes"
	"strconv"
	"math"
)

func main() {
	data := util.ReadInput("input.txt")
	instructs := parse(data)
	run(instructs)
}

func parse(b []byte) []point{
	var moves []point
	lines := bytes.Split(b, []byte("\n"))
	
	for i := range lines{
		
		var p point

		//skip white space and numbers

		if lines[i][0] == byte('R'){
			p = point{toint(lines[i] ),0}
		}
		if lines[i][0] == byte('L'){
			p = point{toint(lines[i] )*-1,0}
		}

		if lines[i][0] == byte('U'){
			p = point{0, toint(lines[i] )}
		}
		if lines[i][0] == byte('D'){
			p = point{0, toint(lines[i] )*-1}
		}

		moves = append(moves, p)
	}
	return moves
}

type point struct {
	x,y int
}

func toint(b []byte)int{
	//strip \n
	num := bytes.Split(b, []byte(" "))
	n,_ := strconv.Atoi(string(num[1]))
	return n
}

func run(steps []point){
	head := point{0,0}
	tail := point{0,0}
	visited := make(map[point]int)
	for _, p := range steps{
		head.x  += p.x
		head.y += p.y
		fmt.Println(distance(head,tail))
		visited[head]++
	}

}

func distance(p,q point)float64{
	return math.Sqrt(float64((p.x - q.x)*(p.x -q.x)) + float64((p.y - q.y)*(p.y-q.y)))
}
