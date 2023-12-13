package main

import (
	"strconv"
	"strings"

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
func run(part2 bool, input string) any {


	data := strings.Split(input, "\n")
	timeLine := data[0][strings.Index(data[0],":")+1:]
	distanceLine:= data[1][strings.Index(data[1],":")+1:]

	timeLine = strings.Trim(timeLine," ")
	distanceLine = strings.Trim(distanceLine," ")

	timeArr := strings.Split(timeLine, " ")
	distanceArr := strings.Split(distanceLine, " ")

	var tempDistArr []string
	
	
	for _,dist := range distanceArr{
		// println("dist=",dist)
		dist = strings.Trim(dist," ")
		if dist !=""{
			tempDistArr = append(tempDistArr, dist)
		}
	}
	distanceArr = tempDistArr

	var tempTimeArr []string

	copy(tempTimeArr,timeArr)

	for _,time := range timeArr{
		time = strings.Trim(time," ")
		if time !=""{
			tempTimeArr = append(tempTimeArr, time)
		}
	}
	timeArr = tempTimeArr

	tempDistStr := ""
	for _,num := range distanceArr{
		tempDistStr += num 
	}
	distanceArr = []string{tempDistStr}


	tempNumStr := ""
	for _,num := range timeArr{
		tempNumStr += num 
	}
	timeArr = []string{tempNumStr}
	ans:=1
	for i,time := range timeArr{	
		cnt:=0

		// println("time=",time)
		currMaxDist,err:= strconv.Atoi(distanceArr[i])
		if err == nil{
			for j:=0;j<currMaxDist;j++{
				timeNum,err:= strconv.Atoi(time)
				if err ==nil{
					speed := j
					secRem:=timeNum-j
					if secRem <0 {
						break
					}
					totalDist := speed*secRem
					// println("speed=",speed," secRem=",secRem,"totalDist=",totalDist)
					if totalDist > currMaxDist{
						cnt++
					}
				}
			}
		}
		ans *=cnt
		// println("Cnt",cnt)
	}

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	return ans
}
