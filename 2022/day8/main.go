package main

import (
	"bytes"
	"fmt"
	"github.com/njr8392/aoc/util"
)

//the data in the input is a square
func main() {
	data := util.ReadInput("input.txt")
	grid := bytes.Split(data, []byte("\n"))
	fmt.Println(numvis(grid))
}

func numvis(d [][]byte) int {
	var vis int
	col := make([]byte, len(d))
	//left & right
	for  i := range d{
		vis += horz(d[i])
	}

	//height relative to top and bottom
	for i := range d{
		for j := range d{
			col[j] = d[j][i]
		}
		fmt.Println(col)
		vis += horz(d[i])

	}

	return vis + (4* len(d))
}

func horz(b []byte) int {
	var vis int
	curtop := b[0]
	for i := 1; i < len(b); i++ {
		if b[i] > curtop {
			vis++
			curtop = b[i]
		}
	}
	curtop = b[len(b)-1]
	for i := len(b) - 2; i > len(b); i-- {
		if b[i] > curtop {
			vis++
			curtop = b[i]
		}
	}
	return vis
}

func checkheight(b byte, height byte) (int, byte) {

	if b > height {
		return 1, b
	}
	return 0, height
}
