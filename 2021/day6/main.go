package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)


func main() {
	var fish []int
	data := ReadInput("input.txt")
	d := bytes.TrimSpace(data)
	chars := bytes.Split(d, []byte{','})
	for _, c := range chars {
		fish = append(fish, ByteToInt(c))
	}
	ans := part1(fish,80)
	ans2 := part1(fish,256)
	fmt.Println(len(ans))
	fmt.Println(len(ans2))

}

//super slow O(n^2). Can I do better?
func part1(nums []int, days int) []int {
	for i := 0; i < days; i++ {
		for j := 0; j < len(nums); j++ {
			if nums[j] == 0 {
				nums[j] = 7
				nums = append(nums, 9)
			}
		}
		for i := range nums {
			nums[i] = nums[i] - 1
		}
	}
	return nums
}
func ReadInput(f string) []byte {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	stats, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, stats.Size())
	_, err = file.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	return buf
}
func ByteToInt(b []byte) int {
	var x int
	for _, num := range b {
		x = 10*x + int(num-'0')
	}
	return x
}
