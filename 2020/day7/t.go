package main

import (
	"fmt"
	"strings"
	"strconv"
	"regexp"
	)

type bag struct {
	name          string
	children      map[string]*bag
	childrenCount map[string]int
}

func (b *bag) contain(name string) bool {
	for childName, child := range b.children {
		if childName == name || child.contain(name) {
			return true
		}
	}
	return false
}

func (b *bag) countBags() int {
	currentCount := 1
	for k, count := range b.childrenCount {
		currentCount += count * b.children[k].countBags()
	}
	return currentCount
}

func parse(s string) []*bag {
	bags := []*bag(nil)
	index := make(map[string]*bag)

	for _, line := range strings.Split(s, "\n") {
		name := regexp.MustCompile(`(.+) bags contain`).FindAllStringSubmatch(line, -1)[0][1]
		b := &bag{
			name:          name,
			children:      make(map[string]*bag),
			childrenCount: make(map[string]int),
		}
		bags = append(bags, b)
		index[b.name] = b

		for _, matches := range regexp.MustCompile(`(\d+) (.+?) bags?`).FindAllStringSubmatch(line, -1) {
			childName := matches[2]
			childCount, _ := strconv.Atoi(matches[1])

			b.children[childName] = nil
			b.childrenCount[childName] = childCount
		}
	}
	for _, b := range bags {
		for k := range b.children {
			b.children[k] = index[k]
		}
	}
	return bags
}

func main(){
	bags := parse(data)

	part1, part2 := 0, 0
	for _, b := range bags {
		if b.name == "shiny gold" {
			part2 = b.countBags() - 1 // remove shiny gold from the count
		}
		if b.contain("shiny gold") {
			part1++
		}
	}

	fmt.Printf("part 1: %d\npart 2: %d", part1, part2)
}
