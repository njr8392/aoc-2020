package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type seatInfo struct {
	info       []string
	errRate    int
}

func (s seatInfo) GenerateRanges() map[int]bool {
	index := make(map[int]bool)
	for _, st := range s.info {
		nums := strings.Split(st, "-")
		n0 := toInt(nums[0])
		n1 := toInt(nums[1])

		for j := 0; j < n1-n0; j++ {
			index[n0+j+1] = true
		}
	}
	return index
}
func toInt(s string) int {
	n, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println(err)
	}
	return n
}


func main() {
	var sInfo seatInfo
	//split on blank lines
	s := strings.Split(data, "\n\n")

	//get all numbers in a range
	ranges := regexp.MustCompile(`(\d+-\d+)`).FindAllString(s[0], -1)

	for _, nums := range ranges {
		sInfo.info = append(sInfo.info, nums)
	}

	rangeMap := sInfo.GenerateRanges()

	//get nearby tickets
	nearby := regexp.MustCompile(`\d+`).FindAllString(s[2], -1)

	for i, ticket := range nearby {
		numtick := toInt(ticket)

		if !rangeMap[numtick] {
			sInfo.errRate += numtick
			nearby[i] = "0"

		}
	}

	fmt.Println(sInfo.errRate)
}
