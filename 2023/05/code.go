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
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	var seeds []int

	seedToSoil := make(map[int]int)
	soilToFert :=make(map[int]int)
	fertToWater :=make(map[int]int)
	waterToLight :=make(map[int]int)
	lightToTemp :=make(map[int]int)
	tempToHum :=make(map[int]int)
	humToLoc := make( map[int]int)


	data := strings.Split(input, "\n")

	for i,line := range data{
		if strings.Contains(line,"seeds"){
			seedsLine:=strings.Split(line, ":")
			seedsLine[1] = strings.Trim(seedsLine[1]," ")
			seedLine:=strings.Split(seedsLine[1]," ")
			for _,seed := range seedLine{
				num,err := strconv.Atoi(seed)
				if err ==nil {
					seeds = append(seeds,num)
				}
			}
		}
		// fmt.Printf("%v",seeds)
		if strings.Contains(line,"seed-to-soil "){
			// seedToSoilLines := strings.Split(line,":")
			for{
				i++
				currLine:=data[i]
				if currLine==""{
					break
				}
				currNums:= strings.Split(currLine," ")
				srcStart,_:= strconv.Atoi(currNums[1])
				destStart,_:= strconv.Atoi(currNums[0])
				rnge,_:= strconv.Atoi(currNums[2])
				for k:=0;k<rnge;k++ {
					seedToSoil[srcStart]=destStart
					srcStart++
					destStart++
				}
				
			}
		}
		if strings.Contains(line,"soil-to-fertilizer "){
			// seedToSoilLines := strings.Split(line,":")
			for{
				i++
				currLine:=data[i]
				if currLine==""{
					break
				}
				currNums:= strings.Split(currLine," ")
				srcStart,_:= strconv.Atoi(currNums[1])
				destStart,_:= strconv.Atoi(currNums[0])
				rnge,_:= strconv.Atoi(currNums[2])
				for k:=0;k<rnge;k++ {
					soilToFert[srcStart]=destStart
					srcStart++
					destStart++
				}
				
			}
		}
		if strings.Contains(line,"fertilizer-to-water "){
			// seedToSoilLines := strings.Split(line,":")
			for{
				i++
				currLine:=data[i]
				if currLine==""{
					break
				}
				currNums:= strings.Split(currLine," ")
				srcStart,_:= strconv.Atoi(currNums[1])
				destStart,_:= strconv.Atoi(currNums[0])
				rnge,_:= strconv.Atoi(currNums[2])
				for k:=0;k<rnge;k++ {
					fertToWater[srcStart]=destStart
					srcStart++
					destStart++
				}
			}
		}
		if strings.Contains(line,"water-to-light "){
			// seedToSoilLines := strings.Split(line,":")
			for{
				i++
				currLine:=data[i]
				if currLine==""{
					break
				}
				currNums:= strings.Split(currLine," ")
				srcStart,_:= strconv.Atoi(currNums[1])
				destStart,_:= strconv.Atoi(currNums[0])
				rnge,_:= strconv.Atoi(currNums[2])
				for k:=0;k<rnge;k++ {
					waterToLight[srcStart]=destStart
					srcStart++
					destStart++
				}
			}
		}
		if strings.Contains(line,"light-to-temperature "){
			// seedToSoilLines := strings.Split(line,":")
			for{
				i++
				currLine:=data[i]
				if currLine==""{
					break
				}
				currNums:= strings.Split(currLine," ")
				srcStart,_:= strconv.Atoi(currNums[1])
				destStart,_:= strconv.Atoi(currNums[0])
				rnge,_:= strconv.Atoi(currNums[2])
				for k:=0;k<rnge;k++ {
					lightToTemp[srcStart]=destStart
					srcStart++
					destStart++
				}
			}
		}
		if strings.Contains(line,"temperature-to-humidity "){
			// seedToSoilLines := strings.Split(line,":")
			for{
				i++
				currLine:=data[i]
				if currLine==""{
					break
				}
				currNums:= strings.Split(currLine," ")
				srcStart,_:= strconv.Atoi(currNums[1])
				destStart,_:= strconv.Atoi(currNums[0])
				rnge,_:= strconv.Atoi(currNums[2])
				for k:=0;k<rnge;k++ {
					tempToHum[srcStart]=destStart
					srcStart++
					destStart++
				}
			}
		}
		if strings.Contains(line,"humidity-to-location "){
			// seedToSoilLines := strings.Split(line,":")
			for{
				i++
				currLine:=data[i]
				if currLine==""{
					break
				}
				currNums:= strings.Split(currLine," ")
				srcStart,_:= strconv.Atoi(currNums[1])
				destStart,_:= strconv.Atoi(currNums[0])
				rnge,_:= strconv.Atoi(currNums[2])
				for k:=0;k<rnge;k++ {
					humToLoc[srcStart]=destStart
					srcStart++
					destStart++
				}
			}
		}
	}
	
	fmt.Printf("%v",seedToSoil)
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	return 42
}