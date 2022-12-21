package main

import (
	"bytes"
	"fmt"
	"github.com/njr8392/aoc/util"
)

//win key
var m = map[byte]byte{
	byte('A'): byte('Y'),
	byte('B'): byte('Z'),
	byte('C'): byte('X'),
}

//loss key
var l = map[byte]byte{
	byte('A'): byte('Z'),
	byte('B'): byte('X'),
	byte('C'): byte('Y'),
}

//score key
var key = map[byte]int{
	byte('X'): 1,
	byte('Y'): 2,
	byte('Z'): 3,
}

func main() {

	data := util.ReadInput("input.txt")
	rounds := bytes.Split(data, []byte("\n"))

	fmt.Println(p1(rounds))
	fmt.Println(p2(rounds))
}

func p2(data [][]byte) int {
	total := 0
	for _, r := range data {
		//2 to ignore whitespace
		if r[2] == 'Z' {
			total += key[m[r[0]]] + 6
		} else if r[2] == 'Y' {
			total += key[r[0]+byte('X'-'A')] + 3
		} else {

			total += key[l[r[0]]]

		}
	}
	return total
}

func p1(data [][]byte) int {
	total := 0
	for _, r := range data {
		//2 to ignore whitespace

		if m[r[0]] == r[2] {
			total += key[r[2]] + 6

		} else if r[2]-r[0] == byte('X'-'A') {
			total += key[r[2]] + 3

		} else {
			total += key[r[2]]
		}
	}
	return total
}
