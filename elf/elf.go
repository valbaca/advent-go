// Package elf: elves are Santa's little helpers!
// Utility functions, in particular short, unsafe versions of functions useful for advent solns
package elf

import (
	"strconv"
	"strings"
)

// c/o https://stackoverflow.com/a/6878625
// I've flipped the order; it shows how they're derived

const MinUint = 0        // 000...
const MaxUint = ^uint(0) // 111....

const MaxInt = int(MaxUint >> 1) // 0111....
const MinInt = -MaxInt - 1       // 1000..

// UnsafeAtoi is strconv.Atoi that simply panics on any error
func UnsafeAtoi(s string) int {
	if out, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return out
	}
}

func UnsafeAtoUint16(s string) uint16 {
	if out, err := strconv.ParseUint(s, 10, 16); err != nil {
		panic(err)
	} else {
		return uint16(out)
	}
}

func Max(ints ...int) int {
	max := ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] > max {
			max = ints[i]
		}
	}
	return max
}

func ParseInt(s string) int {
	if strings.HasSuffix(s, ",") {
		s = s[:len(s)-1]
	}
	return UnsafeAtoi(s)
}

func Sum(a []int) int {
	s := 0
	for _, n := range a {
		s += n
	}
	return s
}

func Dedupe(a []int) []int {
	set := make(map[int]bool)
	for _, n := range a {
		set[n] = true
	}
	// in The Future, may be able to just: maps.Keys(set) but avoiding /x/exp imports...for now.
	deduped := make([]int, len(set))
	i := 0
	for key := range set {
		deduped[i] = key
		i++
	}
	return deduped
}