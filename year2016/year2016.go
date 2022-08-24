package year2016

import (
	"fmt"
	"valbaca.com/advent/year2016/day1"
	"valbaca.com/advent/year2016/day2"
	"valbaca.com/advent/year2016/day3"
	"valbaca.com/advent/year2016/day4"
	"valbaca.com/advent/year2016/day5"
)

func ExecuteYear2016(day int, input string) {
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
	}
}