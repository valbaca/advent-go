package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"valbaca.com/advent2015/day1"
	"valbaca.com/advent2015/day10"
	"valbaca.com/advent2015/day11"
	"valbaca.com/advent2015/day12"
	"valbaca.com/advent2015/day13"
	"valbaca.com/advent2015/day14"
	"valbaca.com/advent2015/day15"
	"valbaca.com/advent2015/day16"
	"valbaca.com/advent2015/day17"
	"valbaca.com/advent2015/day18"
	"valbaca.com/advent2015/day19"
	"valbaca.com/advent2015/day2"
	"valbaca.com/advent2015/day3"
	"valbaca.com/advent2015/day4"
	"valbaca.com/advent2015/day5"
	"valbaca.com/advent2015/day6"
	"valbaca.com/advent2015/day7"
	"valbaca.com/advent2015/day8"
	"valbaca.com/advent2015/day9"
)

func main() {
	day := getDay()
	if day == 0 {
		for i := 1; i <= 19; i++ {
			start := time.Now()

			fmt.Printf("Day %d: ", i)
			executeDay(i)

			elapsed := time.Since(start)
			fmt.Printf("took %.4f\n\n", elapsed.Seconds())
		}
	} else {
		executeDay(day)
	}
}

func executeDay(day int) {
	input := readInputFile(day)
	input = strings.TrimSpace(input)
	switch day {
	case 1:
		fmt.Println(day1.Part1(input), day1.Part2(input))
	case 2:
		fmt.Println(day2.Part1(input), day2.Part2(input))
	case 3:
		fmt.Println(day3.Part1(input), day3.Part2(input))
	case 4:
		fmt.Println(day4.Part1(input), day4.Part2(input))
	case 5:
		fmt.Println(day5.Part1(input), day5.Part2(input))
	case 6:
		fmt.Println(day6.Part1(input), day6.Part2(input))
	case 7:
		fmt.Println(day7.Part1(input, "a"), day7.Part2(input, "a"))
	case 8:
		fmt.Println(day8.Part1(input), day8.Part2(input))
	case 9:
		fmt.Println(day9.BothParts(input))
	case 10:
		fmt.Println(day10.Part1(input), day10.Part2(input))
	case 11:
		fmt.Println(day11.Part1(input), day11.Part2(input))
	case 12:
		fmt.Println(day12.Part1(input), day12.Part2(input))
	case 13:
		fmt.Println(day13.Part1(input), day13.Part2(input))
	case 14:
		fmt.Println(day14.Part1(input, 2503), day14.Part2(input, 2503))
	case 15:
		fmt.Println(day15.Part1(input), day15.Part2(input))
	case 16:
		fmt.Println(day16.Part1(input), day16.Part2(input))
	case 17:
		fmt.Println(day17.Part1(input, 150), day17.Part2(input, 150))
	case 18:
		fmt.Println(day18.Part1(input, 100, 100), day18.Part2(input, 100, 100))
	case 19:
		molecule := "CRnCaSiRnBSiRnFArTiBPTiTiBFArPBCaSiThSiRnTiBPBPMgArCaSiRnTiMgArCaSiThCaSiRnFArRnSiRnFArTiTiBFArCaCaSiRnSiThCaCaSiRnMgArFYSiRnFYCaFArSiThCaSiThPBPTiMgArCaPRnSiAlArPBCaCaSiRnFYSiThCaRnFArArCaCaSiRnPBSiRnFArMgYCaCaCaCaSiThCaCaSiAlArCaCaSiRnPBSiAlArBCaCaCaCaSiThCaPBSiThPBPBCaSiRnFYFArSiThCaSiRnFArBCaCaSiRnFYFArSiThCaPBSiThCaSiRnPMgArRnFArPTiBCaPRnFArCaCaCaCaSiRnCaCaSiRnFYFArFArBCaSiThFArThSiThSiRnTiRnPMgArFArCaSiThCaPBCaSiRnBFArCaCaPRnCaCaPMgArSiRnFYFArCaSiThRnPBPMgAr"
		fmt.Println(day19.Part1(input, molecule), day19.Part2(input, molecule))
	}
}

func getDay() int {
	if len(os.Args) < 2 {
		return 0
	}
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return 0
	}
	return day
}

func readInputFile(day int) string {
	name := fmt.Sprintf("./inputs/day%v.txt", day)
	out, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(out)
}
