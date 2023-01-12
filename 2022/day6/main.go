package main

import (
	"fmt"
	"github.com/njr8392/aoc/util"
)

func main() {
	data := util.ReadInput("input.txt")
	fmt.Println(p1(data))
}

//find first instance of 4 unique characters
func p1(b []byte) int {
	buf := make([]byte, 4)
	for i := range b{
		copy(buf, b[i:i+4])
		fmt.Println(string(buf))
		m := make(map[byte]bool)
		for _, char := range buf {
			m[char] = true
		}

		if len(m) == 4 {
			return i + 4
		}
	}
	return 0
}
