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
	dirsize(*start)
	less := finddirs(*start)
	ttl :=0
	for _, d := range less{
		ttl += d.size
	}
	fmt.Println(ttl)
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
func finddirs(root *dir) []*dir{
	var d []*dir	
	if root.size <= 100000{
		d = append(d, root)
	}

	for _, x := range root.child{
		d = append(d, finddirs(x)...)
	}
	return d
}

func dirsize(root *dir) {
		for _, c := range root.f {
			root.size += c.size
		}
	for _, r := range root.child {
		dirsize(r)
		root.size += r.size
	}
}

