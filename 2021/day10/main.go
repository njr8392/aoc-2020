package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

var valid map[byte]byte = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func main() {
	data := ReadInput("input.txt")
	corrupt := ReadLines(data)
	score := map[byte]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	total := 0
	for key, val := range corrupt {
		total += score[key] * val
	}
	fmt.Println(total)
}

//Read each line and check if it contains valid parenthesis. keep pushing items onto the stack until you reach a closing bracket
//when you reach a closing bracket and that char is not the valid opening char of the top most item on the stack then it is
//invalid
func ReadLines(data []byte) map[byte]int {
	buf := bytes.NewBuffer(data)
	m := make(map[byte]int)

	i := 0
	for {
		var stack []byte
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
		}
		i++
	}
	return m
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
