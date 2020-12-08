package main

import (
	"fmt"
	"strings"
)

type bag struct {
	num  int
	name string
}

func main() {
	m := make(map[string][]string)

	lines := strings.Split(data, "\n")
	for _, line := range lines {
		var childbags []string
		var keystr string
		//split on contains and strip period at the end of the sentence
		bag := strings.Split(line[:len(line)-2], "contain")

		//get the first two words "adj and color"
		colorKey := strings.Split(bag[0], " ")
		colorKey = colorKey[:2]
		keystr = colorKey[0] + " " + colorKey[1]
		children := strings.Split(bag[1], ",")

		for _, desc := range children {
			var childstr string
			newdesc := strings.TrimSpace(desc)
			tmp := strings.Split(newdesc, " ")
			//[ 5 striped tomato bag]
			childstr = tmp[1] + " " + tmp[2]
			childbags = append(childbags, childstr)
		}
		m[keystr] = append(m[keystr], childbags...)

	}

	fmt.Println(m)
	s := visit(m)
	fmt.Println(len(s))
}

func visit(m map[string][]string) []string {
	var contains []string
	seen := make(map[string]bool)
	var visitAll func(bags []string)

	visitAll = func(bags []string) {
		for _, bag := range bags {
			if !seen[bag] {
				seen[bag] = true
				visitAll(m[bag])
				contains = append(contains, bag)
			}
		}
	}

	for key, val := range m {
		for _, child := range val {
			//going to the key i just checked, need to find where that key is a value and work from there
			if strings.Contains(child, "shiny gold") {
				visitAll(m[key])
			}
		}
	}

	return contains
}
