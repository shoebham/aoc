package main

import (
	"math"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)
type Pair struct {
    a, b int
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
	// when you're ready to do part 2, remove this "not implemented" block
	var seeds []int

	seedToSoil :=  make([][]Pair,0)
	soilToFert := make([][]Pair,0)
	fertToWater := make([][]Pair,0)
	waterToLight := make([][]Pair,0)
	lightToTemp := make([][]Pair,0)
	tempToHum := make([][]Pair,0)
	humToLoc :=  make([][]Pair,0)


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
			k:=0
			for{
				i++
				if i>= len(data){break}
				currLine:=data[i]
				if currLine==""{
					break
				}
				currNums:= strings.Split(currLine," ")
				srcStart,_:= strconv.Atoi(currNums[1])
				destStart,_:= strconv.Atoi(currNums[0])
				rnge,_:= strconv.Atoi(currNums[2])
				for len(seedToSoil) <= k {
					seedToSoil = append(seedToSoil, make([]Pair, 2)) // Initialize inner slice with length 2
				}

				// Ensure inner slices have enough capacity to store elements
				if cap(seedToSoil[k]) < 2 {
					newSlice := make([]Pair, 2, 2*cap(seedToSoil[k])) // Double the capacity if needed
					copy(newSlice, seedToSoil[k]) // Copy existing data
					seedToSoil[k] = newSlice // Assign the new slice
				}
				seedToSoil[k][0].a = srcStart
				seedToSoil[k][0].b = destStart
				seedToSoil[k][1].a = srcStart+rnge-1
				seedToSoil[k][1].b = destStart+rnge-1
				k++
				
			}
		}
		if strings.Contains(line,"soil-to-fertilizer "){
			// seedToSoilLines := strings.Split(line,":")
			k:=0
			for{
				i++
				if i>= len(data){break}
				currLine:=data[i]
				if currLine==""{
					break
				}
				currNums:= strings.Split(currLine," ")
				srcStart,_:= strconv.Atoi(currNums[1])
				destStart,_:= strconv.Atoi(currNums[0])
				rnge,_:= strconv.Atoi(currNums[2])
				for len(soilToFert) <= k {
					soilToFert = append(soilToFert, make([]Pair, 2)) // Initialize inner slice with length 2
				}

				// Ensure inner slices have enough capacity to store elements
				if cap(soilToFert[k]) < 2 {
					newSlice := make([]Pair, 2, 2*cap(soilToFert[k])) // Double the capacity if needed
					copy(newSlice, soilToFert[k]) // Copy existing data
					soilToFert[k] = newSlice // Assign the new slice
				}
				soilToFert[k][0].a = srcStart
				soilToFert[k][0].b = destStart
				soilToFert[k][1].a = srcStart+rnge-1
				soilToFert[k][1].b = destStart+rnge-1
				k++
			}
		}
		if strings.Contains(line,"fertilizer-to-water "){
			k:=0
			// seedToSoilLines := strings.Split(line,":")
			for{
				i++
				if i>= len(data){break}
				currLine:=data[i]
				if currLine==""{
					break
				}
				currNums:= strings.Split(currLine," ")
				srcStart,_:= strconv.Atoi(currNums[1])
				destStart,_:= strconv.Atoi(currNums[0])
				rnge,_:= strconv.Atoi(currNums[2])
			
				
				for len(fertToWater) <= k {
					fertToWater = append(fertToWater, make([]Pair, 2)) // Initialize inner slice with length 2
				}

				// Ensure inner slices have enough capacity to store elements
				if cap(fertToWater[k]) < 2 {
					newSlice := make([]Pair, 2, 2*cap(fertToWater[k])) // Double the capacity if needed
					copy(newSlice, fertToWater[k]) // Copy existing data
					fertToWater[k] = newSlice // Assign the new slice
				}
				fertToWater[k][0].a = srcStart
				fertToWater[k][0].b = destStart
				fertToWater[k][1].a = srcStart+rnge-1
				fertToWater[k][1].b = destStart+rnge-1
				k++
			}
		}
		if strings.Contains(line,"water-to-light "){
			// seedToSoilLines := strings.Split(line,":")
			k:=0
			for{
				i++
				if i>= len(data){break}
				currLine:=data[i]
				if currLine==""{
					break
				}
				currNums:= strings.Split(currLine," ")
				srcStart,_:= strconv.Atoi(currNums[1])
				destStart,_:= strconv.Atoi(currNums[0])
				rnge,_:= strconv.Atoi(currNums[2])
				
				for len(waterToLight) <= k {
					waterToLight = append(waterToLight, make([]Pair, 2)) // Initialize inner slice with length 2
				}

				// Ensure inner slices have enough capacity to store elements
				if cap(waterToLight[k]) < 2 {
					newSlice := make([]Pair, 2, 2*cap(waterToLight[k])) // Double the capacity if needed
					copy(newSlice, waterToLight[k]) // Copy existing data
					waterToLight[k] = newSlice // Assign the new slice
				}
				
				
				waterToLight[k][0].a = srcStart
				waterToLight[k][0].b = destStart
				waterToLight[k][1].a = srcStart+rnge-1
				waterToLight[k][1].b = destStart+rnge-1
				k++
			}
		}
		if strings.Contains(line,"light-to-temperature "){
			// seedToSoilLines := strings.Split(line,":")
			k:=0
			for{
				i++
				if i>= len(data){break}
				currLine:=data[i]
				if currLine==""{
					break
				}
				currNums:= strings.Split(currLine," ")
				srcStart,_:= strconv.Atoi(currNums[1])
				destStart,_:= strconv.Atoi(currNums[0])
				rnge,_:= strconv.Atoi(currNums[2])



				for len(lightToTemp) <= k {
					lightToTemp = append(lightToTemp, make([]Pair, 2)) // Initialize inner slice with length 2
				}

				// Ensure inner slices have enough capacity to store elements
				if cap(lightToTemp[k]) < 2 {
					newSlice := make([]Pair, 2, 2*cap(lightToTemp[k])) // Double the capacity if needed
					copy(newSlice, lightToTemp[k]) // Copy existing data
					lightToTemp[k] = newSlice // Assign the new slice
				}
				
				
			
				lightToTemp[k][0].a = srcStart
				lightToTemp[k][0].b = destStart
				lightToTemp[k][1].a = srcStart+rnge-1
				lightToTemp[k][1].b = destStart+rnge-1
				k++
			}
		}
		if strings.Contains(line,"temperature-to-humidity "){
			// seedToSoilLines := strings.Split(line,":")
			k:=0
			for{
				i++
				if i>= len(data){break}
				currLine:=data[i]
				if currLine==""{
					break
				}
				currNums:= strings.Split(currLine," ")
				srcStart,_:= strconv.Atoi(currNums[1])
				destStart,_:= strconv.Atoi(currNums[0])
				rnge,_:= strconv.Atoi(currNums[2])
				


				for len(tempToHum) <= k {
					tempToHum = append(tempToHum, make([]Pair, 2)) // Initialize inner slice with length 2
				}

				// Ensure inner slices have enough capacity to store elements
				if cap(tempToHum[k]) < 2 {
					newSlice := make([]Pair, 2, 2*cap(tempToHum[k])) // Double the capacity if needed
					copy(newSlice, tempToHum[k]) // Copy existing data
					tempToHum[k] = newSlice // Assign the new slice
				}

				tempToHum[k][0].a = srcStart
				tempToHum[k][0].b = destStart
				tempToHum[k][1].a = srcStart+rnge-1
				tempToHum[k][1].b = destStart+rnge-1
				k++
			}
		}
		if strings.Contains(line,"humidity-to-location "){
			// seedToSoilLines := strings.Split(line,":")
			k:=0
			for{
				i++
				if i>= len(data){break}
				currLine:=data[i]
				if currLine==""{
					break
				}
				currNums:= strings.Split(currLine," ")
				srcStart,_:= strconv.Atoi(currNums[1])
				destStart,_:= strconv.Atoi(currNums[0])
				rnge,_:= strconv.Atoi(currNums[2])
			

				for len(humToLoc) <= k {
					humToLoc = append(humToLoc, make([]Pair, 2)) // Initialize inner slice with length 2
				}

				// Ensure inner slices have enough capacity to store elements
				if cap(humToLoc[k]) < 2 {
					newSlice := make([]Pair, 2, 2*cap(humToLoc[k])) // Double the capacity if needed
					copy(newSlice, humToLoc[k]) // Copy existing data
					humToLoc[k] = newSlice // Assign the new slice
				}


				humToLoc[k][0].a = srcStart
				humToLoc[k][0].b = destStart
				humToLoc[k][1].a = srcStart+rnge-1
				humToLoc[k][1].b = destStart+rnge-1
				k++
			}
		}
	}

	// fmt.Printf("seedToSoil: %v\n",seedToSoil)
	// fmt.Printf("soilToFert: %v\n",soilToFert)
	// fmt.Printf("fertToWater: %v\n",fertToWater)
	// fmt.Printf("waterToLight: %v\n",waterToLight)
	// fmt.Printf("lightToTemp: %v\n",lightToTemp)
	// fmt.Printf("tempToHum: %v\n",tempToHum)
	// fmt.Printf("humToLoc: %v\n",humToLoc)
	ans:=math.MaxInt32
	for i:=0;i<len(seeds);i+=2{
		seed:=seeds[i]
		low:=seed
		high:=seed+seeds[i+1]
		if i==len(seeds)-1{break}
		for{
			// println("low=",low,"high=",high)
			if (low>high){
				break
			}
			seed=(low+high)/2
			// println(seed)
			var value  int
			value = helper(seedToSoil,seed)
			// println("seedtosoil",value)

			value = helper(soilToFert,value)
			// println("soilToFert",value)

			value = helper(fertToWater,value)
			// println("fertToWater",value)

			value = helper(waterToLight,value)
			// println("waterToLight",value)

			value = helper(lightToTemp,value)
			// println("lightToTemp",value)

			value = helper(tempToHum,value)
			// println("tempToHum",value)

			value = helper(humToLoc,value)
			// println("humToLoc",value)

			// println("seed=",seed)
			// println("value=",value)
			if value < ans{
				high=seed-1
				ans = min(ans,value)
			}else if value>ans{
				low = seed+1
			}
		}
		// println("SEED: ",seed)

	// println("---------")

		
	}
	// num := seeds[0]

	// ans:=humToLoc[tempToHum[lightToTemp[waterToLight[fertToWater[soilToFert[seedToSoil[num]]]]]]]
	// println("Answer:",ans)
	// fmt.Printf("%v",seedToSoil)
	if part2 {

	}
	// solve part 1 here
	return ans
}

func helper(parr [][]Pair, seed int) int{
	var value int
	for _,p := range parr{
		if seed >= p[0].a && seed <= p[1].a{
				value = p[0].b+(seed-p[0].a)
				break
		}else {
			value = seed
		}
	}
	return value
}
func getValueOrDefault(m map[int]int, key int) int {
	if val, ok := m[key]; ok {
		return val
	}
	return key
}