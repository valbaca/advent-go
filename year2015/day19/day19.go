package day19

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Part1(in, s string) string {
	reps := ParseInput(in)
	return fmt.Sprintf("%v", CountDistincts(reps, s))
}

func Part2(in, s string) string {
	reps := ParseInput(in)
	return fmt.Sprintf("%v", MinPath(reps, "e", s))
}

type Replacements map[string][]string

func ParseInput(in string) Replacements {
	reps := Replacements{}

	for _, line := range strings.Split(in, "\n") {
		fields := strings.Fields(line)
		old, new := fields[0], fields[2]
		if val, ok := reps[old]; ok {
			reps[old] = append(val, new)
		} else {
			reps[old] = []string{new}
		}
	}
	return reps
}

func CountDistincts(reps Replacements, s string) int {
	set := GetNewUniques(reps, s)
	return len(set)
}

func GetNewUniques(reps Replacements, s string) map[string]bool {
	set := map[string]bool{}
	for a, list := range reps {
		for _, b := range list {
			// a => b
			si := 0 // start index
			for si >= 0 && si < len(s) {
				idx := strings.Index(s[si:], a)
				if idx == -1 {
					break
				} else {
					idx += si
				}
				end := idx + len(a)
				new := s[0:idx] + b + s[end:]
				set[new] = true
				si = end
			}
		}
	}
	return set
}

func MinPath(reps Replacements, org, tgt string) int {
	red := buildRed(reps)
	atom := tgt
	count := 0
	for atom != "e" {
		shrunk := false
		for _, pair := range red {
			lng, short := pair.lng, pair.short
			temp := strings.Replace(atom, lng, short /* all */, -1)
			if temp != atom {
				count += strings.Count(atom, lng)
				atom = temp
				shrunk = true
				break
			}
		}
		if !shrunk { // reset, shuffle, try again (this is like bogosort...but it works)
			atom = tgt
			count = 0
			shuffle(&red)
		}
	}
	return count
}

type Pair struct{ lng, short string }

func buildRed(reps Replacements) []Pair {
	var red []Pair
	for short, lngs := range reps {
		for _, lng := range lngs {
			red = append(red, Pair{lng, short})
		}
	}
	return red
}

func shuffle(lst *[]Pair) {
	// https://yourbasic.org/golang/shuffle-slice-array/
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*lst), func(i, j int) {
		(*lst)[i], (*lst)[j] = (*lst)[j], (*lst)[i]
	})
}
