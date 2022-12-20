package main

import (
	"fmt"
	"io"
	"os"
	"bytes"
	"strconv"
	"sort"
)

func ReadInput(file string) []byte {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	s, _ := f.Stat()
	buf := make([]byte, s.Size()-1)
	_, er := f.Read(buf)
	if er != nil && er != io.EOF {
		fmt.Println(er)
		return nil
	}
	return buf

}
func main() {
	txt := ReadInput("inout.txt")

	cals := bytes.Split(txt, []byte("\n\n"))
	ans := max(cals)
	fmt.Printf("p1 = %d\n", ans)
}


func max(food [][]byte)int{
	var top []int
	for i := range food{
		nums := bytes.Split(food[i], []byte("\n"))
		sum :=0
		for j := range nums{
			cal,_ := strconv.Atoi(string(nums[j]))
			sum += cal
		}
		top = append(top, sum)
	}
	sort.Ints(top)
	total := top[len(top)-1] + top[len(top)-2] + top[len(top)-3]
	return total
}

