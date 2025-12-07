package day7

import (
	"fmt"

	stl "aoc/stl"
)

func Solution() {
	ip := stl.ReadFile("input.txt")
	// ip := stl.ReadFile("example.txt")

	splits := countSplits(ip)
	fmt.Printf("Tachyon beam split: %d times\n", splits)
}

func countSplits(grid []string) (splits int) {
	numRows := len(grid)

	tachyonBeams := [][]int{{70, 0}} // beamCol: startingRow

	for len(tachyonBeams) > 0 {
		// there might be overlapping beams
		next := make(map[int]int)
		for _, cr := range tachyonBeams {
			col, row := cr[0], cr[1]

			if numRows <= row+1 {
				continue
			}

			if grid[row][col] == '^' {
				splits++

				next[col+1] = row + 1
				next[col-1] = row + 1
			} else {
				next[col] = row + 1
			}
		}

		tachyonBeams = [][]int{}
		for c, r := range next {
			tachyonBeams = append(tachyonBeams, []int{c, r})
		}
	}
	return splits
}
