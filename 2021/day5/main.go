package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type vector struct {
	x int
	y int
}

func SetVector(x, y int) vector {
	v := vector{x: x, y: y}
	return v
}

func main() {
	set := make(map[vector]int)
	data := ReadInput("input.txt")
	lines := bytes.Split(data, []byte("\n"))
	for _, l := range lines {
		points := bytes.Split(l, []byte(" -> "))

		// split into vectors v1 = x1,y1 and v2 = x2,y2
		v1 := bytes.Split(points[0], []byte(","))
		v2 := bytes.Split(points[1], []byte(","))
		fmt.Println(v2)

		//set bytes to int to compare vals
		x0 := ByteToInt(v1[0])
		x1 := ByteToInt(v2[0])
		y0 := ByteToInt(v1[1])
		y1 := ByteToInt(v2[1])

		//only care where x or y equal eachother
		if x0 == x1 {
			max, min := Max(y0, y1)
			for i := min; i <= max; i++ {
				point := vector{x: x0, y: i}
				set[point]++
			}
		}

		if y0 == y1 {
			max, min := Max(x0, x1)
			for i := min; i <= max; i++ {
				point := vector{x: i, y: y0}
				set[point]++
			}
		}
	}
	ans := CountSeen(set)
	fmt.Println(ans)
}
func CountSeen(m map[vector]int) int {
	count := 0
	for _, num := range m {
		if num >= 2 {
			count++
		}
	}
	return count
}

func Max(x, y int) (int, int) {
	if x > y {
		return x, y
	}
	return y, x
}

func ByteToInt(b []byte) int {
	var n int
	for _, byt := range b {
		n = n*10 + int(byt-'0')
	}
	return n
}

func ReadInput(file string) []byte {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	stat, err := f.Stat()
	buf := make([]byte, stat.Size())
	_, err = f.Read(buf)
	if err != io.EOF && err != nil {
		log.Fatal(err)
	}
	//gets rid of the last \n char
	return buf[:len(buf)-1]
}
