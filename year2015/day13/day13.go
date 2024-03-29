// Package day13
// TIL: passing a slice to a function really passes the 'header' for the slice
//
//	This is effectively like passing a pointer.
//	So for recursive functions like this, it's best to defensively clone it
//
// TIL: clone := append(orig[:0:0], orig...) // efficient clone
package day13

import (
	"strconv"
	"strings"
	"valbaca.com/advent/elf"
)

type name string

type feelings map[name]int

type people map[name]feelings

func Part1(in string) string {
	people := parseInput(in)
	return strconv.Itoa(findOptimal(people))
}

func Part2(in string) string {
	people := parseInput(in)
	people.addNilSelf()
	return strconv.Itoa(findOptimal(people))
}

func parseInput(in string) people {
	p := people{}
	sp := strings.Split(in, "\n")
	for _, line := range sp {
		if line == "" {
			continue
		}
		a, diff, b := parseLine(line)
		p.addFeeling(name(a), name(b), diff)
	}
	return p
}

// Store how a feels about sitting next to b
func (p people) addFeeling(a, b name, diff int) {
	curr, ok := p[a]
	if !ok {
		curr = feelings{}
	}
	curr[b] = diff
	p[a] = curr
}

// Get how a feels about sitting next to b
func (p people) getFeeling(a, b name) int {
	return (p[a])[b]
}

func parseLine(line string) (string, int, string) {
	// Alice would gain 54 happiness units by sitting next to Bob.
	// 0     1     2    3  4         5     6  7       8    9  10
	sp := strings.Split(line, " ")
	if len(sp) != 11 {
		panic("invalid line given:" + line)
	}
	a, gainOrLose, diffStr, b := sp[0], sp[2], sp[3], sp[10]
	pos := false
	if gainOrLose == "gain" {
		pos = true
	}
	diff := elf.UnsafeAtoi(diffStr)
	if !pos {
		diff = -1 * diff
	}
	b = b[:len(b)-1] // remove the trailing period
	return a, diff, b
}

func findOptimal(p people) int {
	return findRecur([]name{}, p)
}

func findRecur(s []name, p people) int {
	seated := append(s[:0:0], s...) // clone
	if len(seated) == len(p) {
		return sumHappiness(seated, p)
	}
	toAttemptInSeat := []name{}
	for name := range p {
		if !contains(seated, name) {
			toAttemptInSeat = append(toAttemptInSeat, name)
		}
	}
	max := elf.MinInt
	n := len(toAttemptInSeat)
	maxes := make(chan int, n)
	for _, attempt := range toAttemptInSeat {
		go subFindRecur(seated, attempt, p, maxes)
	}
	for i := 0; i < n; i++ {
		m := <-maxes
		if m > max {
			max = m
		}
	}
	return max
}

func subFindRecur(seated []name, attempt name, p people, results chan int) {
	attemptedSeating := append(seated, attempt)
	recurResult := findRecur(attemptedSeating, p)
	results <- recurResult
}

func contains(names []name, query name) bool {
	for _, n := range names {
		if n == query {
			return true
		}
	}
	return false
}

func sumHappiness(seatings []name, p people) int {
	sum := 0
	for i := 0; i < len(seatings)-1; i++ {
		a, b := seatings[i], seatings[i+1]
		sum += p.getFeeling(a, b)
		sum += p.getFeeling(b, a)
	}
	first, last := seatings[0], seatings[len(seatings)-1]
	sum += p.getFeeling(first, last)
	sum += p.getFeeling(last, first)
	return sum
}

func (p people) addNilSelf() {
	names := make([]name, len(p))
	for k := range p {
		names = append(names, k)
	}
	nilName := name("nil")
	for _, name := range names {
		p.addFeeling(name, nilName, 0)
		p.addFeeling(nilName, name, 0)
	}
}
