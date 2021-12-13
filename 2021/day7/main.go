package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var pos []int
	data := ReadInput("input.txt")
	trimed := bytes.TrimSpace(data)
	removecomas := bytes.Split(trimed, []byte{','})
	for _, b := range removecomas {
		pos = append(pos, ByteToInt(b))
	}
	min := MinFuel(pos)
	fmt.Println(min)
}
func MinFuel(nums []int) int {
	var minfuel int = 0x0fffffffffffffff
	max := Max(nums)
	for i := 0; i <= max; i++ {
		tmp := 0
		for _, pos := range nums {
			tmp += abs(pos - i)
		}
		if tmp < minfuel {
			minfuel = tmp
		}
	}
	return minfuel
}
func Max(nums []int) int {
	var max int
	for _, x := range nums {
		if x > max {
			max = x
		}
	}
	return max
}
func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
func ByteToInt(b []byte) int {
	var num int
	for _, x := range b {
		num = num*10 + int(x-'0')
	}
	return num
}

func ReadInput(file string) []byte {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	stats, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, stats.Size())
	_, err = f.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	return buf
}
