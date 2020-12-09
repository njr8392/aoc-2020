package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseint(s string) (int64, error) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return -1, err
	}
	return num, nil
}


func read(s []string) (int64, error) {
	var linenum int64
	var accum int64
	seen := make(map[int64]bool)

	for !seen[linenum] {
		instruction := strings.Split(s[linenum], " ")
		action := instruction[0]
		number, err := parseint(instruction[1])
		seen[linenum] = true
		//fmt.Printf("action:%s\t number:%d\n", action, number)

		if err != nil {
			return -1, fmt.Errorf("Error parsing int64: %s", err)
		}

		switch action {

		case "nop":
			linenum++

		case "acc":
			accum += number
			linenum++

		case "jmp":
			linenum += number
		}

	}

	return accum, nil
}

func main() {
	set := strings.Split(data, "\n")
	part1, err := read(set)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(part1)
}
