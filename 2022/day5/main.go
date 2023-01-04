package main

import(
	"github.com/njr8392/aoc/util"
	"fmt"
	"bytes"
	)

func main(){
	f := util.ReadInput("input.txt")

	split := bytes.Split(f, []byte("\n\n"))
	stack := split[0]
	instruct := split[1]

	board := parseStack(stack)
}

func makeStack()[][]byte{
	var tmp []byte
	stack := make([][]byte, 9)
	for i := range stack{
		stack[i] = append(stack[i], tmp)
	}

	return stack
}

func parseStack(f []byte)[][]byte{
	stack := makeStack()
	count :=0
	for _, val := range f {
		if val == ' '{
			count++
		}
		if val == '\n'{
			count =0
		}
		if val >=  'A' && val <= 'Z'{
			stack[count] = append(stack[count], val)
		}
	}
	return stack
}
