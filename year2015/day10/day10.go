// Package day10
// TIL: two int-to-string conversions: strconv.Itoa(int) and fmt.Sprintf("%v", int)
// TIL: bytes.Buffer is WAY faster than string appending
package day10

import (
	"bytes"
	"fmt"
	"strconv"
)

func Part1(in string) string {
	n := in
	for i := 0; i < 40; i++ {
		n = lookAndSay(n)
	}
	return fmt.Sprintf("%v", len(n))
}

func Part2(in string) string {
	n := in
	for i := 0; i < 50; i++ {
		n = lookAndSay(n)
	}
	return fmt.Sprintf("%v", len(n))
}

func lookAndSay(n string) string {
	// Buffer is WAAAYY faster than strings
	// runtime when from >10mins to 0.4s
	var buffer bytes.Buffer
	runLen := 1
	var curr byte
	for i := 0; i < len(n); i++ {
		r := n[i]
		if i == 0 {
			curr = r
		} else if curr == r {
			runLen++
		} else if curr != r {
			buffer.Write([]byte(strconv.Itoa(runLen)))
			buffer.WriteByte(curr)
			curr = r
			runLen = 1
		}
	}
	buffer.Write([]byte(strconv.Itoa(runLen)))
	buffer.WriteByte(curr)
	return buffer.String()
}
