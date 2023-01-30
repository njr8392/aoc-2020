package main

import (
	"bytes"
	"fmt"
	"github.com/njr8392/aoc/util"
	"strconv"
)
const ROUNDS =20 
type monkey struct {
	n       int
	items   []int
	opstr   string
	opnum   string
	test    int
	iftrue  int
	iffalse int
	inspect int
}

func main() {
	input := util.ReadInput("input.txt")
	data := parse(input)
	p1(data)
	
}

func p1(monks []*monkey){
	for i:=0; i < ROUNDS; i++{
	for _, m := range monks{
	 	for _, item := range m.items{
			m.inspect++
			num, err := strconv.Atoi(m.opnum)
			if err != nil{
				num = item
			}
			switch m.opstr{
				case "+":
					item += num
				case "*":
					item *= num
			}

			item /= 3
			if item % m.test == 0{
				monks[m.iftrue].items = append(monks[m.iftrue].items, item)
			}

		}
	}
	}
	for _, m := range monks{
		fmt.Println(m.inspect)
	}
	fmt.Println(1540*1700)
}

func parse(b []byte) []*monkey {
	var monks []*monkey
	set := bytes.Split(b, []byte("\n\n"))
	for i, mky := range set {
		m := new(monkey)
		m.n = i
		line := bytes.Split(mky, []byte("\n"))
		for j, l := range line {
			switch j {
			case 0:
				continue
			case 1:
				split := bytes.Split(l, []byte(": "))
				nums := bytes.Split(split[1], []byte(", "))
				for _, n := range nums {
					toint, _ := strconv.Atoi(string(n))
					m.items = append(m.items, toint)
				}
			case 2:
				split := bytes.Split(l, []byte("old "))
				m.opstr = string(split[1][0])
				m.opnum = string(split[1][2:])
			case 3:
				split := bytes.Split(l, []byte("by "))
				num, _ := strconv.Atoi(string(split[1]))
				m.test = num
			case 4:
				split := bytes.Split(l, []byte("monkey "))
				num, _ := strconv.Atoi(string(split[1]))
				m.iftrue = num
			case 5:
				split := bytes.Split(l, []byte("monkey "))
				num, _ := strconv.Atoi(string(split[1]))
				m.iffalse = num
			}
		}
		monks = append(monks, m)
	}
	return monks
}

