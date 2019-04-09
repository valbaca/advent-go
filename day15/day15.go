// TIL: how to pass func to func is made much easier with typedef and cleaner than in C
package day15

import (
	"fmt"
	"math/big"
	"strings"

	"valbaca.com/advent2015/utils"
)

const TOTAL_TSP = 100
const TARGET_CALS = 500

func Part1(in string) string {
	parts := ParseInput(in)
	return Optimize(parts, GetScore)
}

func Part2(in string) string {
	parts := ParseInput(in)
	return Optimize(parts, GetScoreCal)
}

func ParseInput(in string) []Part {
	out := []Part{}
	for _, line := range strings.Split(in, "\n") {
		out = append(out, ParseLine(line))
	}
	return out
}

func ParseLine(line string) Part {
	// Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
	// 0             1        2   3          4   5      6  7       8  9        10
	sp := strings.Split(line, " ")
	name := sp[0]
	capacity := int64(utils.ParseInt(sp[2]))
	durability := int64(utils.ParseInt(sp[4]))
	flavor := int64(utils.ParseInt(sp[6]))
	texture := int64(utils.ParseInt(sp[8]))
	calories := int64(utils.ParseInt(sp[10]))
	return Part{name, capacity, durability, flavor, texture, calories}
}

type Part struct {
	name       string
	capacity   int64
	durability int64
	flavor     int64
	texture    int64
	calories   int64
}

//type Drink []Measurements

type Measurements struct {
	tsp  int64
	part Part
}

func Optimize(parts []Part, fn scoreFunc) string {
	drink := []Measurements{}
	return fmt.Sprintf("%v", OptRecur(parts, TOTAL_TSP, drink, fn))
}

type scoreFunc func(d []Measurements) *big.Int

func OptRecur(parts []Part, left int64, drink []Measurements, fn scoreFunc) *big.Int {
	if len(parts) == 1 {
		rest := Measurements{left, parts[0]}
		d := append(drink[:0:0], drink...)
		d = append(d, rest)
		return fn(d)
	}

	max := &big.Int{}
	for i := int64(0); i <= left; i++ {
		m := Measurements{i, parts[0]}
		d := append(drink[:0:0], drink...)
		d = append(d, m)
		score := OptRecur(parts[1:], left-i, d, fn)
		if score.Cmp(max) > 0 {
			max = score
		}
	}
	return max
}

func GetScore(d []Measurements) *big.Int {
	var capSum, durSum, flavSum, texSum int64
	for _, m := range d {
		t, p := m.tsp, m.part
		capSum += t * p.capacity
		durSum += t * p.durability
		flavSum += t * p.flavor
		texSum += t * p.texture
	}
	if capSum <= 0 || durSum <= 0 || flavSum <= 0 || texSum <= 0 {
		return &big.Int{}
	}
	score := &big.Int{}
	score.Mul(big.NewInt(capSum), big.NewInt(durSum))
	score.Mul(score, big.NewInt(flavSum))
	score.Mul(score, big.NewInt(texSum))
	return score
}

func GetScoreCal(d []Measurements) *big.Int {
	var capSum, durSum, flavSum, texSum, calSum int64
	for _, m := range d {
		t, p := m.tsp, m.part
		capSum += t * p.capacity
		durSum += t * p.durability
		flavSum += t * p.flavor
		texSum += t * p.texture
		calSum += t * p.calories
	}
	if calSum != TARGET_CALS {
		return &big.Int{}
	}
	if capSum <= 0 || durSum <= 0 || flavSum <= 0 || texSum <= 0 {
		return &big.Int{}
	}
	score := &big.Int{}
	score.Mul(big.NewInt(capSum), big.NewInt(durSum))
	score.Mul(score, big.NewInt(flavSum))
	score.Mul(score, big.NewInt(texSum))
	return score
}
