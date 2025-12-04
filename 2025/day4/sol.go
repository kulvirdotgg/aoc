package day4

import (
	"fmt"

	stl "aoc/stl"
)

// https://adventofcode.com/2025/day/4
func Solution() {
	// ip := stl.ReadFile("example.txt")
	ip := stl.ReadFile("input.txt")
	pickable := pickableRolls(ip)
	fmt.Printf("Pickable Rolls for 4 or less adjancent: %d\n", pickable)
}

func pickableRolls(grid []string) (pickable int) {
	n, m := len(grid), len(grid[0])

	// Check the 8 adjacent directions of given cell
	// returns the count of adjancent paper rolls
	adjacentRolls := func(r, c int) (adjacent int) {
		// adjancent directions for any given cell
		directions := [][]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}

		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]

			validPosition := nc >= 0 && nr >= 0 && nr < n && nc < m
			if validPosition && grid[nr][nc] == '@' {
				adjacent++
			}
		}
		return adjacent
	}

	for r, row := range grid {
		for c := range row {
			if row[c] == '@' {
				adjacentRolls := adjacentRolls(r, c)
				if adjacentRolls < 4 {
					pickable++
				}
			}
		}
	}
	return pickable
}
