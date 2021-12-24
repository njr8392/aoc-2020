package main

import(
	"bytes"
	"io"
	"github.com/njr8392/aoc/util"
	"fmt"
	)

func main(){
	input := util.ReadInput("input.txt")
	template:= bytes.Split(input, []byte("\n\n"))
	set := ParseInstruct(template[1])
	polymer := IterPolymer(template[0], set, 10)
	p1 := CountChars(polymer)
	fmt.Printf("Ans to part 1 is %d\n", p1)
}
func IterPolymer(poly []byte, set map[string]byte, step int)[]byte{
	for step >0{
		toinsert := make([]byte, len(poly)-1)
		buf := make([]byte, len(poly)*2-1)

		for i:=0; i < len(poly)-1; i++{
			toinsert[i] = set[string(poly[i:i+2])]   // ignore last element in poly
		}
		j :=1
		buf[0] = poly[0]
		for _, x := range toinsert{
			if x == 0{
				fmt.Println("addin zeros", step)
				break
			}
		}
		for i:=1; i < len(poly); i++{
			buf[j] = toinsert[i-1]
			buf[j+1] = poly[i]
			j+=2
		}
		poly = buf
		step--
	}
	return poly
}
//swap end with second to last element in the buf
func swap(b []byte, i int)[]byte{
	b[i], b[i-1] = b[i-1], b[i]
	return b
}

func CountChars(poly []byte)int{
	var max int
	var min int = 1<<63-1
	m := make(map[byte]int)
	for _, b := range poly{
		m[b]++
	}
	fmt.Println(m)
	for _, num := range m{
		max = Max(num, max)
		min = Min(num,min)
	}
	return max-min
}

func ParseInstruct(instruct []byte)map[string]byte{
	buf := bytes.NewBuffer(instruct)
	m := make(map[string]byte)
	for {
		line, err := buf.ReadBytes('\n')
		if err != nil{
			if err == io.EOF{
				break
			}
		}
		rule:= bytes.Split(line[:len(line)-1], []byte(" -> "))
		m[string(rule[0])] = rule[1][0] //just want the byte to insert not the who slice
	}
	return m
}
func Max(x,y int)int{
	if x >y{
		return x 
	}
	return y
}

func Min(x,y int)int{
	if x <y{
		return x
	}
	return y
}
