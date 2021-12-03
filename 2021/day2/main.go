package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ReadInput(file string) []byte {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	info, _ := f.Stat()
	buf := make([]byte, info.Size())
	_, err = f.Read(buf)
	if err != nil && err != io.EOF {
		panic(err)
	}
	return buf
}

func main() {
	aim := 0
	depth := 0
	fwrd := 0
	data := ReadInput("input.txt")
	buf := bytes.NewBuffer(data)
	split := []byte{' '}

	for {
		// include delim in return
		line, err := buf.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		instruct := bytes.Split(line[:len(line)-1], split)
		switch string(instruct[0]) {
		case "forward":
			fwrd += toInt(instruct[1])
			depth += toInt(instruct[1]) * aim

		case "down":
			//comment out aim for part 1's answer
			//				depth += toInt(instruct[1]) //only for p1
			aim += toInt(instruct[1])

		case "up":
			//				depth -= toInt(instruct[1]) // only for part2
			aim -= toInt(instruct[1])

		default:
			fmt.Printf("unrecongizd input %s\n", string(instruct[0]))
		}
	}

	fmt.Printf("Part 1: depth = %d forward = %d ans = %d\n", depth, fwrd, fwrd*depth)
	fmt.Printf("Part2: %d\n", depth*fwrd)

}

func toInt(b []byte) int {
	conv, err := strconv.Atoi(string(b))
	if err != nil {
		panic(err)
	}
	return conv
}
