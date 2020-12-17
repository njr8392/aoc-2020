package main

import (
	"fmt"
)

type elf struct {
	seen map[int]bool
	last map[int][]int
	turn int
}

func (e *elf) Seen(i int) bool {
	if _, ok := e.seen[i]; !ok {
		e.seen[i] = true
		return false
	}
	return true
}

func (e *elf) newNum(key int) int {
	o := e.last[key]
	newnum := o[1] - o[0]
	return newnum
}

func start() *elf {
	m := make(map[int]bool)
	l := make(map[int][]int)
	e := elf{m, l, 1}

	return &e
}
func main() {
	var spoke int
	var spoke1 int
	game := start()
	s := []int{13,0,10,12,1,5,8}

	for i, num := range s {
		game.seen[num] = true
		game.last[num] = append(game.last[num], i+1)
		game.turn++
	}

	spoke = 0
	for i := len(s) + 1; i < 30000000; i++ {
		game.last[spoke] = append(game.last[spoke], i)

		if len(game.last[spoke]) > 2 {
			game.last[spoke] = game.last[spoke][1:]
		}

		if game.Seen(spoke) {
			spoke = game.newNum(spoke)
		} else {
			spoke = 0
		}
		
		if i == 2019{
			spoke1 = spoke
		}
		game.turn++
	}
	
	fmt.Printf("Solution to part 1: %d Solution to part 2: %d\n", spoke1, spoke)
}
