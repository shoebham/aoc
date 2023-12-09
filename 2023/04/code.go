package main

import (
	"math"
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
	data := strings.Split(input, "\n")
	sum:=0.0
	for _,line := range data{
		cnt:=0.0
		cards := line[(strings.Index(line,":"))+1:]
		cards = strings.TrimSpace(cards)
		w1:= strings.Split(cards, "|")[0]
		h1:=strings.Split(cards,"|")[1]
		haveMap := make(map[int]int)

		w1 = strings.TrimSpace(w1)
		h1 = strings.TrimSpace(h1)
		// println(w1)
		// println(h1)
		winning := strings.Split(w1," ")
		have := strings.Split(h1," ")
		// fmt.Printf("%v",have)

		for _,num :=range have{
			if num==" "{
				continue
			}
			currNum,err := strconv.Atoi(num)
			if err == nil{
				// println("in have num",num)
				haveMap[currNum]++
			}
		}
		for _,num := range winning{
			if num==" "{
				continue
			}
			currNum,err := strconv.Atoi(num)
			if err == nil{
				if haveMap[currNum]!=0 {
					// println("in winning num",num)
					// println("havemap[num]=",haveMap[currNum])
				cnt++
				}
			}
			
		}
		// println(int(cnt))
		if cnt>0{
			sum+=math.Pow(2,cnt-1)
		}
		// println("sum=",int(sum))
		// println("-----")

	}
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	return sum
}
