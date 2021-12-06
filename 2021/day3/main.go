package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

const NSIZE = 12
const FSIZE = 1000

func main() {
	data := ReadInput("input.txt")
	ans := solve(data)
	fmt.Println(ans)

}

func solve(data []byte) int {
	var gam int
	var ep int 
	for i := 0; i < NSIZE; i++ {
		count := 0
		buf := bytes.NewBuffer(data)
		for {
			line, err := buf.ReadBytes('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			// only concerned with 0-11
			if line[len(line)-i-2] == '1' {
				count++
			}
		}
		if count > FSIZE/2 {
			gam |= (1 << i)
		}else {
			ep |= (1<< i)
		}
	}
	return gam * ep
}

func ReadInput(file string) []byte {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	info, _ := f.Stat()
	buf := make([]byte, info.Size())
	_, err = f.Read(buf)
	if err != nil && err != io.EOF {
		panic(err)
	}
	return buf
}
