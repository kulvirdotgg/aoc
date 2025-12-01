package day1

import (
	"fmt"
	"strconv"

	stl "aoc/stl"
)

// https://adventofcode.com/2025/day/1
func Solution() {
	ip := stl.ReadFile("input.txt")
	zeroCounts := crackPassword(ip)

	fmt.Println(zeroCounts)
}

func crackPassword(doc []string) (zeroCount int) {
	// starting position of safe tick is 50
	// and, There are total of 100 ticks on the safe, i.e. 0 - 99
	var dialPos, ticksOnSafe int64 = 50, 100

	for _, rotation := range doc {
		dir, ticks := rotation[0], rotation[1:]

		ticksCount, err := strconv.ParseInt(ticks, 10, 64)
		if err != nil {
			panic(err)
		}

		switch dir {
		case 'L':
			dialPos = (dialPos - ticksCount) % ticksOnSafe
		case 'R':
			dialPos = (dialPos + ticksCount) % ticksOnSafe
		}

		if dialPos == 0 {
			zeroCount++
		}
	}

	return zeroCount
}
