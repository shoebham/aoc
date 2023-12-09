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
	copy := make([]int,len(data)+1)
	cntarr  :=make([]int,len(data)+1)
	sum:=0.0
	part2ans:=0
	for i,line := range data{
		i=i+1
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
		if cnt>0{
			sum+=math.Pow(2,cnt-1)
		}
		cntarr[i]=int(cnt)
		for j:=i+1;j<=i+int(cnt);j++{
			// println("j=",j,"cnt=",int(cnt))
			copy[j]++

		}

		// for j:=i+1;j<=i+int(cnt);j++{
		// 	// println("j=",j+1,"cnt=",int(cnt))
		// 	for z:=0;z<copy[j];z++{
		// 		copy[j]++
		// 	}
		// }

		// println()
		// println("cnt=",int(cnt))

	}
	for k,num :=range copy{
		// println("k=",k,"num=",num,"cntarr=",cntarr[k])
		for z:=k+1;z<=k+cntarr[k];z++{
			// println("z=",z)
			copy[z]+=num
		}
	}
	// println("--------after---------")
	for _,num :=range copy{
		part2ans+=num
		// println("k=",k,"num=",num,"cntarr=",cntarr[k])
	}
	part2ans+=len(copy)-1
	if part2 {
		return part2ans
	}
	// solve part 1 here
	return sum
}

// 4, 4, 4
// 2, 4, 2
// 2