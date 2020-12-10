package main

import(
	"fmt"
	"strings"
	"strconv"
	)

func parse(s string)int{
	xmas := strings.Split(s, "\n")
	nums := convStrArray(xmas)	

	for i:= 25; i<len(nums); i++{
		if !sumPrev25(nums, nums[i], i){
			
			return nums[i]
		}
	}
	return -1
}

func convStrArray(s []string)[]int{
	nums := make([]int, len(s))

	for i, str := range s{
		num, err := strconv.Atoi(str)

		if err != nil{
			fmt.Println(err)
		}

		nums[i] = num
	}
	return nums
}
func sumPrev25(nums []int, num int, index int)bool{
	last25 := make([]int, 25)
	
	copy(last25, nums[index-25:index+1])

	for i, n := range last25{
		for j, m :=  range last25{
			if i != j{
				tmp := n + m
				if tmp == num{
					return true
				}
			}
		}
	}

	return false
}

func main(){
part1 := parse(data)
fmt.Printf("solution to part 1 is %d\n", part1)
}
