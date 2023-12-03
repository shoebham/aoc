package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and examle input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	data:= strings.Split(input, "\n")

	sum:=0
	for i:=1;i<6;i++{
		sum+=i
	}
	// example only 12 red cubes, 13 green cubes, and 14 blue cubes
	// part 1 12 red cubes, 13 green cubes, and 14 blue cubes
	part2ans:=0
	for i,line := range data{
		maxr,maxg,maxb := -1,-1,-1
		// sets := make([][]string,100)	
		line = line[(strings.Index(line,":"))+1:]
		sets := strings.Split(line,";")
		// fmt.Println("i=",i+1)	
		// fmt.Println("line",line)
		// var red,green,blue int[]
		for _, set := range sets{
			colors := strings.Split(set,",")
			ok:=true
			for _,color := range colors{
				color = strings.Trim(color," ")
				// fmt.Println("Color",color)

				red:=strings.Index(color,"red")
				green:=strings.Index(color,"green")
				blue:=strings.Index(color,"blue")

				// fmt.Printf("red:%d, blue:%d, green:%d\n",red,blue,green)
				
				if(red!=-1){
					red,_=strconv.Atoi(color[:strings.Index(color," ")]) 
					maxr= max(red,maxr)
				}
				if(green!=-1){
					green,_ = strconv.Atoi(color[:strings.Index(color," ")])
					maxg = max(green,maxg)
				}
				if(blue!=-1){
					blue,_=strconv.Atoi(color[:strings.Index(color," ")])
					maxb = max(maxb,blue)
				}
				if !part2{
					if(red>12||green>13||blue>14){

						sum-=(i+1)
						ok=false
						break
					}
				}
			}
			if !ok {
				break
			}
			// fmt.Println("set ended")
		}
		fmt.Println("red",maxr,"blue",maxb,"green",maxg)
		part2ans+=(maxr*maxg*maxb)
	}
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return part2ans
	}
	// solve part 1 here
	return sum
}
