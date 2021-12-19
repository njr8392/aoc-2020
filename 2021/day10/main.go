package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

var valid map[byte]byte = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func main() {
	data := ReadInput("input.txt")
	corrupt, p2_score := ReadLines(data)
	score := map[byte]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	//part1
	total := 0
	for key, val := range corrupt {
		total += score[key] * val
	}
	fmt.Printf("answer to part1 is %d\n", total)

	//part2
	sort.Ints(p2_score)
	fmt.Printf("answer to part2 is %d\n", p2_score[len(p2_score)/2])
}

//Read each line and check if it contains valid parenthesis. keep pushing items onto the stack until you reach a closing bracket
//when you reach a closing bracket and that char is not the valid opening char of the top most item on the stack then it is
//invalid. For part 2 if the line is valid but just incomplete we will calculate the score with the remaining items in the stack
func ReadLines(data []byte) (map[byte]int, []int) {
	buf := bytes.NewBuffer(data)
	m := make(map[byte]int)
	var p2_score []int
	i := 0
	for {
		var stack []byte
		incomplete := false
		line, err := buf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		//remove the newline char
		for pos, char := range line[:len(line)-1] {
			if char == '(' || char == '{' || char == '<' || char == '[' {
				stack = append(stack, char)
			}

			//err check in case the first char is not an opening bracket
			if top, err := GetTop(stack); err == nil {
				if char == ')' || char == ']' || char == '}' || char == '>' {
					if valid[top] != char {

						fmt.Printf("expected %c but got %c on line %d at posistion %d\n", valid[top], char, i, pos)
						m[char]++
						break
					}
					stack = stack[:len(stack)-1]
				}
			}
			if pos == len(line[:len(line)-1])-1 {
				incomplete = true
			}
		}

		if incomplete {
			p2_score = append(p2_score, IncompleteScore(stack))
		}
		i++
	}
	return m, p2_score
}
func IncompleteScore(stack []byte) int {
	var score map[byte]int = map[byte]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	total := 0
	i := len(stack) - 1
	for i >= 0 {
		total *= 5
		total += score[valid[stack[i]]]
		i--
	}
	return total
}

func GetTop(b []byte) (byte, error) {
	if len(b) != 0 {
		return b[len(b)-1], nil
	}
	return 0, fmt.Errorf("stack is empty")
}

func ReadInput(file string) []byte {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	buf := make([]byte, info.Size())
	_, err = f.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
		return nil
	}
	return buf

}
