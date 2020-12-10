package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func findEncoding(s string) (int, int) {
	xmas := strings.Split(s, "\n")
	nums := convStrArray(xmas)
	p1 := 0
	p2 := 0
	index := 0

	for i := 25; i < len(nums); i++ {
		if !sumPrev25(nums, nums[i], i) {

			p1 = nums[i]
			index = i
			break
		}
	}

	p2 = part2(nums[:index], p1)

	return p1, p2
}

//part2 two find continguous set of at least 2 numbers
func part2(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		sum := 0
		set := nums[i:]
	findsum:
		for j, num := range set {
			sum += num

			if sum == val {
				desired := set[:j+1]
				sort.Ints(desired)

				return desired[0] + desired[len(desired)-1]

			}

			if sum > val {
				break findsum
			}
		}
	}

	return -1
}
func convStrArray(s []string) []int {
	nums := make([]int, len(s))

	for i, str := range s {
		num, err := strconv.Atoi(str)

		if err != nil {
			fmt.Println(err)
		}

		nums[i] = num
	}
	return nums
}
func sumPrev25(nums []int, num int, index int) bool {
	last25 := make([]int, 25)

	copy(last25, nums[index-25:index+1])

	for i, n := range last25 {
		for j, m := range last25 {
			if i != j {
				tmp := n + m
				if tmp == num {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	part1, part2 := findEncoding(data)
	fmt.Printf("solution to part 1 is %d\nSolution to part 2 is %d\n", part1, part2)
}
