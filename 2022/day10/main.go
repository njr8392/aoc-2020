package main

import (
	"bytes"
	"fmt"
	"github.com/njr8392/aoc/util"
	"strconv"
)

func main() {
	input := util.ReadInput("input.txt")
	data := parse(input)
	s := solution(data)
	fmt.Println(s)
}

func solution(ops []instruct) int {
	var q []int
	cycle := 0
	addr := 1
	sum := 0
	for i := 0; i < len(ops) || len(q) > 0; i++ {
		cycle++
		if len(q) > 0 {
			n := q[0]
			q = q[1:]
			addr += n
		}
		if (cycle+20)%40 == 0 && cycle <= 220 {
			strength := addr * cycle
			sum += strength
		}
		//p2
		if cycle%40 == 1 {
			fmt.Println()
		}
		col := (cycle - 1) % 40
		switch addr {
		case col - 1, col, col + 1:
			fmt.Print("#")
		default:
			fmt.Print(".")
		}
		// end of p2

		if i > len(ops)-1 {
			continue
		}
		if ops[i].op == "noop" {
			q = append(q, 0)
		}
		if ops[i].op == "addx" {
			q = append(q, 0, ops[i].num)
		}
	}
	return sum
}

func parse(b []byte) []instruct {
	data := bytes.Split(b, []byte("\n"))
	var list []instruct
	for _, l := range data {
		line := bytes.Split(l, []byte(" "))

		inst := instruct{op: string(line[0])}
		if len(line) > 1 {
			n, _ := strconv.Atoi(string(line[1]))
			inst.num = n
		}

		list = append(list, inst)
	}
	return list
}

type instruct struct {
	op  string
	num int
}
