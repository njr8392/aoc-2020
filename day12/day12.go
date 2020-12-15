package main

import
	(
	"fmt"
	"strings"
	"strconv"
	)

	
type ship struct{
	posistion rune
	north, south, east, west, deg int
}
func (s *ship)Instruction(instr string){
	dir := instr[0] 
	val := instr[1:]

	 num, err := strconv.Atoi(val) 
	 if err!= nil{
		fmt.Println(err)
	}
	
	switch dir{
		case 'N':
			s.north += num

		case 'S':
			s.south += num

		case 'E':
			s.east += num
		case 'W':
			s.west += num
		case 'F':
			s.Instruction(string(s.posistion) + val)

		default:
			s.Rotate(rune(dir), num)
	}
}

func (s *ship) Rotate(direction rune, degree int){
	switch direction{
		case 'R':
			s.deg += degree		
		
		case 'L':
			s.deg -= degree
	}

	s.Abs()
	s.ChangeDirection()
}

func (s * ship)ChangeDirection(){
	switch s.deg{
		case 0:
			s.posistion = 'N'
		case 90:
			s.posistion = 'E'
		case 180:
			s.posistion = 'S'
		case 270:
			s.posistion = 'W'
	}
}
func (s *ship) Abs(){
	if s.deg >= 360{
		s.deg -= 360 
	} else if s.deg < 0 {
		s.deg += 360
	}else{
		s.deg = s.deg
	}
}
func (s *ship)ManhattanSum()int{
	n := abs(s.north - s.south)
	e := abs(s.east - s.west)
	
	return n + e
}

func abs(i int)int{
	if i < 0 {
		return i * -1
	}
	return i 
}

func main(){
sh := ship{posistion: 'E', deg: 90}
input := strings.Split(data, "\n")
for _, pos := range input{
	sh.Instruction(pos)

}
//part 1
fmt.Println(sh.ManhattanSum())
}
