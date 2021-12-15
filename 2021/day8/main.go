package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	count := 0
	input := ReadInput("input.txt")
	buf := bytes.NewBuffer(input)
	for {
		line, err := buf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		data := bytes.Split(line[:len(line)-1], []byte(" | "))
		output := bytes.Split(data[1], []byte(" "))
		for _, seq := range output {
			if len(seq) == 2 || len(seq) == 3 || len(seq) == 4 || len(seq) == 7 {
				count++
			}
		}
	}
	fmt.Println(count)
}

func ReadInput(file string) []byte {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, info.Size())
	_, err = f.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	return buf
}
