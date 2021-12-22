package main

import (
	"bytes"
	"fmt"
	"github.com/njr8392/aoc/util"
	"io"
	"log"
)

type Node struct {
	name     string
	children map[string]*Node
}

func NewNode(name string) *Node {
	return &Node{name: name, children: make(map[string]*Node)}
}
func main() {
	data := util.ReadInput("input.txt")
	start := SetGraph(data)
	fmt.Println(start)
	seen := visit(start)
	//working. printing correct number of nodes ie 13
	// confirmed it is a directed graph
	// problem is all set up now just need to figure out how to find all the paths
	fmt.Println(len(seen))
	paths := FindPaths(start)
	fmt.Println(len(paths))
}
//DFS to walk all paths
func FindPaths(head *Node) [][]string {
	var find func(*Node, []string, bool) [][]string
	find = func(n *Node, path []string,seen bool) [][]string {
		if n.name == "end" {
			path = append(path, n.name)
			return [][]string{path}
		}
		if IsLower(n.name) {
			for _, small := range path{            
				if small == n.name{
					if small == "start"{        //delete here to 61 for part1
						return [][]string{}
					}
					if seen{
						return [][]string{}
					}
					seen = true
				}
			}
		}
		paths := [][]string{}
		path = append(path, n.name)
		for _, child := range n.children {
			paths = append(paths, find(child, path,seen)...)
		}
		return paths
	}
	return find(head, []string{}, false)
}
func IsLower(s string) bool {
	for _, char := range s {
		if char < 'a' {
			return false
		}
	}
	return true
}

//test for printing all nodes
func visit(n *Node) map[*Node]bool {
	var walk func(n *Node)
	seen := make(map[*Node]bool)

	walk = func(n *Node) {
		for h := n; h != nil; {
			if seen[h] {
				return
			}
			fmt.Printf("Node %p %s: %d\n", h, h.name, len(h.children))
			seen[h] = true
			for _, node := range h.children {
				walk(node)
			}
		}
	}
	walk(n)
	return seen
}
func SetGraph(data []byte) *Node {
	var start *Node
	seen := make(map[string]*Node)
	buf := bytes.NewBuffer(data)

	for {
		line, err := buf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		edge := bytes.Split(line[:len(line)-1], []byte{'-'})
		n1, n2 := NewNode(string(edge[0])), NewNode(string(edge[1]))

		//if we already created these nodes then grab them and update the pointers
		if seen[n1.name] != nil {
			n1 = seen[n1.name]
		}
		if seen[n2.name] != nil {
			n2 = seen[n2.name]
		}

		seen[n1.name] = n1
		seen[n2.name] = n2
		n1.children[n2.name] = n2 // add edges between the nodes
		n2.children[n1.name] = n1

		if n1.name == "start" {
			start = n1
		}
	}
	return start
}
