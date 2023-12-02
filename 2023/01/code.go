package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/jpillora/puzzler/harness/aoc"
)
func part1(data []string){
	sum:=0
	for _,line := range data{
		var first, last string
		isFirst:=true
		for _,char := range line{
			if unicode.IsDigit(char){
				if isFirst {
					first = string(char)
					isFirst=false
				}
				last = string(char)
			}
		}
		num:=first+last
		x,_:=strconv.Atoi(num)
		sum+=x

	}
	fmt.Println(sum)
}
func extractDigitOrWord(line string) (int, bool) {
	switch {
	case line[0] >= '0' && line[0] <= '9':
		return int(line[0] - '0'), true
	case strings.HasPrefix(line, "zero"):
		return 0, true
	case strings.HasPrefix(line, "one"):
		return 1, true
	case strings.HasPrefix(line, "two"):
		return 2, true
	case strings.HasPrefix(line, "three"):
		return 3, true
	case strings.HasPrefix(line, "four"):
		return 4, true
	case strings.HasPrefix(line, "five"):
		return 5, true
	case strings.HasPrefix(line, "six"):
		return 6, true
	case strings.HasPrefix(line, "seven"):
		return 7, true
	case strings.HasPrefix(line, "eight"):
		return 8, true
	case strings.HasPrefix(line, "nine"):
		return 9, true
	default:
		return 0, false
	}
}


func part2(data []string){
	sum:=0
	var first,last int
	for _,line := range data{
		for j,_ := range line{
			num,ok := extractDigitOrWord(line[j:])
			if ok{
				first=num
				break
			}
		}
		for j,_ := range line{
			num,ok := extractDigitOrWord(line[len(line)-j-1:])
			if ok{
				last=num
				break
			}
		}
		fmt.Printf("line:%s, first:%d, last:%d \n",line,first,last)
		num:=first*10+last
		sum+=num
	}
	fmt.Println(sum)
}
func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	data := strings.Split(input, "\n")
	// when you're ready to do part 2, remove this "not implemented" block
	sum:=0
	for _,line := range data{
		var first, last string
		isFirst:=true
		for _,char := range line{
			if unicode.IsDigit(char){
				if isFirst {
					first = string(char)
					isFirst=false
				}
				last = string(char)
			}
		}
		num:=first+last
		x,_:=strconv.Atoi(num)
		sum+=x

	}
	
	if part2 {
		sum:=0
		var first,last int
		for _,line := range data{
			for j,_ := range line{
				num,ok := extractDigitOrWord(line[j:])
				if ok{
					first=num
					break
				}
			}
			for j,_ := range line{
				num,ok := extractDigitOrWord(line[len(line)-j-1:])
				if ok{
					last=num
					break
				}
			}
			// fmt.Printf("line:%s, first:%d, last:%d \n",line,first,last)
			num:=first*10+last
			sum+=num
		}	
		return sum
	}
	// solve part 1 here
	return sum
}
