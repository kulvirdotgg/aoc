package day4

import (
	"fmt"
	"strings"

	stl "aoc/stl"
)

// https://adventofcode.com/2025/day/4
func Solution() {
	ip := stl.ReadFile("input.txt")
	pickable := pickableRolls(ip)
	fmt.Printf("Pickable rolls with less than 4 adjancent rolls: %d\n", pickable)

	totalPicked := pickUntilCant(ip)
	fmt.Printf("Total rolls picked till, cant: %d\n", totalPicked)
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

func pickUntilCant(grid []string) (totalPicked int) {
	n, m := len(grid), len(grid[0])

	adjacentRolls := func(r, c int) (adjacent int) {
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

	// count the paper rolls that are picked
	// also contstruct new grid after picking the paper rolls
	pickAndBuild := func(grid []string) (picked int, gridAfterPicking []string) {
		for r, row := range grid {
			var rowBuilder strings.Builder

			for c := range row {
				if row[c] == '@' {
					adjacentRolls := adjacentRolls(r, c)

					if adjacentRolls < 4 {
						// if a roll is picked mark spot as empty
						rowBuilder.WriteByte('.')
						picked++
					} else {
						rowBuilder.WriteByte(grid[r][c])
					}
				} else {
					rowBuilder.WriteByte(grid[r][c])
				}
			}

			gridAfterPicking = append(gridAfterPicking, rowBuilder.String())
		}
		return picked, gridAfterPicking
	}

	picked := 1<<16 - 1
	// continue picking the rolls
	// unless it is not possible to pick any role
	for picked != 0 {
		currentPicked, gridAfterPicking := pickAndBuild(grid)

		totalPicked += currentPicked
		picked = currentPicked

		// update the grid to new grid after picking the rolls
		grid = gridAfterPicking
	}

	return totalPicked
}
