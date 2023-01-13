package main

import (
	"bytes"
	"fmt"
	"github.com/njr8392/aoc/util"
	"strconv"
	"strings"
)

func main() {
	f := util.ReadInput("input.txt")
	d := new(dir)
	var start **dir = &d
	data := bytes.Split(f, []byte("\n"))
	parse(data, d)
	fmt.Println(**start)
	fmt.Println(sum(*start))
}

type file struct {
	size int
	name string
}

type dir struct {
	parent *dir
	name   string
	child  []*dir
	f      []file
	size   int
}

func parse(data [][]byte, state *dir) {
	for i := 2; i < len(data); i++ {
		if data[i][0] == '$' {
			data[i] = data[i][2:]
		}
		line := strings.Split(string(data[i]), " ")
		if line[0] == "ls" {
			continue
		}
		if line[0] == "dir" {
			d := new(dir)
			d.name = line[1]
			d.parent = state
			state.child = append(state.child, d)
		}
		if line[0] == "cd" {
			if line[1] == ".." {
				state = state.parent
			} else {
				for _, c := range state.child {
					if c.name == line[1] {
						state = c
					}
				}
			}
		}

		if line[0] != "dir" && line[0] != "cd" {
			var tmp file
			num, _ := strconv.Atoi(line[0])
			tmp.size = num
			tmp.name = line[1]
			state.f = append(state.f, tmp)
		}

	}
}

//identify the dir that is >100000 then sum all child directories
func sum(root *dir) int{
	ttl := 0
	if root.child == nil {
		return 0
	}
	for i, c := range root.child {
		sum(c)
		if ttl > 100000 && i == len(root.child) {
			return ttl
		}
		for _, s := range c.f {
			ttl += s.size
			fmt.Println(ttl)
		}
	}
	return ttl
}

func dirsize(root *dir) {
	if root.child == nil {
		return
	}
	for _, r := range root.child {
		dirsize(r)
		tmp :=0
		for _, c := range r.f {
			tmp += c.size
		}
		root.size += tmp
	}
}

func dfs(root *dir)int{
	sum :=0
	if root.size <= 100000{
		sum += root.size
	}
	if root.child == nil {
		return 0
	}
	for _, r := range root.child{
		dfs(r)
	}
	return sum
}
