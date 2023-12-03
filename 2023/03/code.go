package main

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func allPeriods(str string) bool {
	for _, char := range str {
		if char != '.' {
			return false
		}
	}
	return true
}

func run(part2 bool, input string) any {
	data := strings.Split(input, "\n")
	// var nums []int
	for i := range data {
    data[i] += "."
	}
	sum:=0
	for i,line :=range data{

		var num string 
		for j,char := range line{
			if !unicode.IsDigit(char){
				currNum,err := strconv.Atoi(num)
				if err == nil{
					sum+=currNum
				// middle one
					if i>0 && i < len(data)-1{
						var up,left,down,right,tl,tr,br,bl string
						right = string(data[i][j])

						up    = data[i-1][j-len(num):j]
						down  = data[i+1][j-len(num):j]
						if j-len(num)>0{
							left  = string(data[i][j-len(num)-1])
							tl = string(data[i-1][j-len(num)-1])
							bl = string(data[i+1][j-len(num)-1])
						}
						// println(currNum)
						// println(up)
						// println(down)
						// println(left)
						// println(right)
							
						if(j<len(line)-1){
							tr = string(data[i-1][j])
							br = string(data[i+1][j])
						}

						if allPeriods(up)&&
							allPeriods(down)&&
							allPeriods(left)&&
							allPeriods(right)&&
							allPeriods(tl)&&
							allPeriods(tr)&&
							allPeriods(bl)&&
							allPeriods(br){
								// println("sum-=num",currNum)
								sum-=currNum
							}
					}else if i==0{
						var left,down,right,br,bl string
						right = string(data[i][j])
						down  = data[i+1][j-len(num):j]
						if j-len(num)>0{
							left  = string(data[i][j-len(num)-1])
							bl = string(data[i+1][j-len(num)-1])
						}
						if(j<len(line)-1){
							br = string(data[i+1][j])
						}
							// println(up)
						// println(currNum)
						// println(down)
						// println(left)
						// println(right)
						if allPeriods(down)&&
							allPeriods(left)&&
							allPeriods(right)&&
							allPeriods(bl)&&
							allPeriods(br){
								// println("sum-=num",currNum)
								sum-=currNum
							}
					}else if i==len(data)-1{

						var up,left,right,tl,tr string
						right = string(data[i][j])

						up    = data[i-1][j-len(num):j]
						if j-len(num)>0{
							left  = string(data[i][j-len(num)-1])
							tl = string(data[i-1][j-len(num)-1])
						
						}
						// println(up)
						// println(down)
						// println(left)
						// println(right)
							
						if(j<len(line)-1){
							tr = string(data[i-1][j])
						}

						if allPeriods(up)&&
							allPeriods(left)&&
							allPeriods(right)&&
							allPeriods(tl)&&
							allPeriods(tr){
								// println("sum-=num",currNum)
								sum-=currNum
							}
					}
				}
				num=""
			}else{
				num += string(char)
			}
			
		}
	}
	// fmt.Printf("nums %v",nums)
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	return sum
}
