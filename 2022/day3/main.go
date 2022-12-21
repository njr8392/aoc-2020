package main

import (
	"bytes"
	"fmt"
	"github.com/njr8392/aoc/util"
)

func main() {
	data := util.ReadInput("input.txt")
	bags := bytes.Split(data, []byte("\n"))
	var common []byte
	p1 := SumPrioity(Split(bags))
	fmt.Println(p1)

	//part2  // double check to identify that is the val is in each row
	for i:=0; i<len(bags); i+=3{
		//find common in group
		m := make(map[byte]bool)
		group := i+2
		for i <= group{
		 	for j := range bags[i]{
				if !m[bags[i][j]{
					m[bags[i][j]] = true
				}
			}
			group--
		}
		
	}


}

//find common element in each list -- part 1
func Split(data [][]byte)[]byte{
	var common []byte
	for _, b := range data {
		mid := len(b) / 2
		comp1 := b[:mid]
		comp2 := b[mid:]
		m := make(map[byte]bool)

		for i := range comp1 {
			m[comp1[i]] = true
		}

		for j := range comp2 {
			if m[comp2[j]] {
				common = append(common, comp2[j])
				break
			}
		}
	}
}

// find sum of priority of each element
//Lowercase item types a through z have priorities 1 through 26.
//Uppercase item types A through Z have priorities 27 through 52.
func SumPriority( els []byte)int{
	sum := 0

	for _, c := range els {
		if c <= 'Z' {
			sum += int(c) - 'A' + 27
			fmt.Println(string(c), int(c) -'A'+27)
		} else {
			sum += int(c) - 'a' + 1
			fmt.Println(string(c), int(c) -'a'+1)
		}

	}
}
