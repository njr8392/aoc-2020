package main

import(
	"os"
	"fmt"
	"io"
	"bytes"
	"strconv"
	)


func ReadInput(file string)[]byte{
	f, err := os.Open(file)
	if err != nil{
		panic(err)
	}
	defer f.Close()
	buf := make([]byte, 10000) // too lazy to call os.Stat to get the len of the file
	_, err = f.Read(buf)
	if err != nil && err != io.EOF{
		panic(err)
	}
	return buf
}

func remove(b []byte)[]byte{
	var buf []byte
	for _, char := range b{
		if char != '\n'{
			buf = append(buf, char)
		}
	}
	return buf
}


func main(){
	var nums []int
	data := ReadInput("./input.txt")
	buf := bytes.NewBuffer(data)
	for {
		num, err := buf.ReadBytes('\n')
		if err == io.EOF{
			break
		}
		n, err := strconv.Atoi(string(num[:len(num)-1]))
		if err != nil{
			panic(err)
		}

		nums = append(nums, n)

	}
	count := NumIncreasing(nums)
	p2 := ThreeMeasure(nums)
	fmt.Println(len(nums))
	fmt.Println(count)
	fmt.Println(p2)
}	

func NumIncreasing(n []int)int{
	count :=0
	prev := n[0]
	for _, num := range n[1:]{
		if num > prev{
			count++
		}
		prev = num 
	}
	return count
}
	
func ThreeMeasure(n []int)int{
	count :=0
	prev := n[0] + n[1] + n[2]
	for i, num := range n[3:]{
		if i +2 > len(n)-1{
			return count
		}
		cur := num + n[i+1] + n[i+2]
		if cur > prev{
			count++
		}
		prev = cur
	}
	return count
}
