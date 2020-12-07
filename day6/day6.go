package main

import (
	"fmt"
	"strings"
)

func main() {
	p1 := 0
	p2 := 0
	t1 := strings.Split(data, "\n\n")

	for _, s := range t1 {
		m := make(map[rune]int)

		people := strings.Split(s, "\n")
		for _, person := range people {
			for _, ans := range person {
					m[ans]++
					
					if m[ans] == len(people){
						p2++
					}
			}
		}
		p1 += len(m)

	}
	fmt.Printf("Answer to part 1 is %d\nAnswer to part 2 is %d\n", p1, p2)
}
