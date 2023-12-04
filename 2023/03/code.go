package main

import (
	"regexp"
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
func extractNum( index int,input string) string {
	// Check if the index is within bounds
	if index < 0 || index >= len(input) || !unicode.IsDigit(rune(input[index])) {
		return "Invalid index or not a digit"
	}

	// Traverse forward to find the end of the number
	end := index
	for end < len(input) && unicode.IsDigit(rune(input[end])) {
		end++
	}

	// Traverse backward to find the start of the number
	start := index
	for start >= 0 && unicode.IsDigit(rune(input[start])) {
		start--
	}

	// Adjust start and end positions to get the whole number
	start++
	if start < 0 {
		start = 0
	}
	if end > len(input) {
		end = len(input)
	}

	return input[start:end]
}

var re = regexp.MustCompile(`\d+`)
func runPart2(data []string)(int){

	ans:=0
	for i,line :=range data{
		for j,char :=range line{

			if char == '*'{
				gears:=[]int{}

				// left
				if j>0 && unicode.IsDigit(rune(line[j-1])){
					// println("left")
					nums:=re.FindAllString(line[:j],-1)
					num,_ := strconv.Atoi(nums[len(nums)-1])	
					gears = append(gears, num)
				}
				//right
				if j<len(line)-1&&unicode.IsDigit(rune(line[j+1])){
					// println("right")
					nums:=re.FindString(line[j+1:])
					num,_ := strconv.Atoi(nums)
					gears = append(gears, num)
				}

				if i>0 {
					// println("up")
					numIdxs := re.FindAllStringIndex(data[i-1],-1)
					// fmt.Println(data[i-1])
					// fmt.Printf("%v\n",numIdxs)
					for _,idx := range numIdxs{
						if (j >= idx[0] && j<= idx[1]) || (j+1)==idx[0]{
							// println(line[idx[0]:idx[1]])
							num,_ := strconv.Atoi(data[i-1][idx[0]:idx[1]])
							gears = append(gears, num)
						}
					}
				}
				if i<len(data)-1{
					// println("down")
					numIdxs := re.FindAllStringIndex(data[i+1],-1)
					for _,idx := range numIdxs{
						if (j >= idx[0] && j<= idx[1]) || (j+1)==idx[0]{
							num,_ := strconv.Atoi(data[i+1][idx[0]:idx[1]])
							gears = append(gears, num)
						}
					}
				}
				// fmt.Printf("gears:%v\n",gears)

				if len(gears)>1{
					ans+=(gears[0]*gears[1])
				}
			}
		}
	}
	// println(ans)
	return ans
}
func part1(data []string)(int){
	var arr []string

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
						println(currNum)
						println(up)
						println(down)
						println(left)
						println(right)

						if(j<len(line)-1){
							tr = string(data[i-1][j])
							br = string(data[i+1][j])
						}
						if(
							strings.Contains(up,"*")||
							strings.Contains(down,"*")||
							strings.Contains(left,"*")||
							strings.Contains(right,"*")||
							strings.Contains(bl,"*")||
							strings.Contains(br,"*")||
							strings.Contains(tl,"*")||
							strings.Contains(tr,"*")){
								arr = append(arr, num)
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
	return sum
}
func run(part2 bool, input string) any {

	data := strings.Split(input, "\n")
	
	// fmt.Printf("nums %v",nums)
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return runPart2(data)
	}
	// solve part 1 here
	return part1(data)
}
