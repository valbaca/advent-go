// Package day2
// Elves are wrapping boxes with WxLxH and need to calculate wrapping paper
// square area and ribbon length
// TIL: pros/cons of go...
//
//	pros: multi-value returns, strings.Split(s, pattern)
//	cons: no ternary operator, awkward string<->int conversions
package day2

import (
	"fmt"
	"strconv"
	"strings"
	"valbaca.com/advent/elf"
)

// Part1 wrapping paper is straightforward. Most of the trouble is converting
// between ints and strings
func Part1(in string) string {
	ss := strings.Split(in, "\n")
	var sum int
	for _, s := range ss {
		if s == "" {
			continue
		}
		a, b, c := toInts(s)
		surface := getSurface(a, b, c)
		slack := getSlack(a, b, c)
		sum += surface + slack
	}
	return strconv.Itoa(sum)
}
func getSurface(a, b, c int) int {
	return 2*a*b + 2*b*c + 2*a*c
}

func getSlack(a, b, c int) int {
	min, mid := getMinAndMid(a, b, c)
	return min * mid
}

// Part2 is similarly easy; again, converting is most of the code
func Part2(in string) string {
	ss := strings.Split(in, "\n")
	var sum int
	for _, s := range ss {
		if s == "" {
			continue
		}
		a, b, c := toInts(s)
		sum += getRibbon(a, b, c)
		sum += getBow(a, b, c)
	}
	return strconv.Itoa(sum)
}

func getRibbon(a, b, c int) int {
	min, mid := getMinAndMid(a, b, c)
	return 2*min + 2*mid
}

func getBow(a, b, c int) int {
	return a * b * c
}

func toInts(line string) (int, int, int) {
	ss := strings.Split(line, "x")
	if len(ss) < 3 {
		panic(fmt.Sprintf("couldn't split %s", line))
	}
	return elf.UnsafeAtoi(ss[0]), elf.UnsafeAtoi(ss[1]), elf.UnsafeAtoi(ss[2])
}

func getMinAndMid(a, b, c int) (int, int) {
	if a >= b && a >= c {
		return b, c
	} else if b >= a && b >= c {
		return a, c
	} else {
		return a, b
	}
}
