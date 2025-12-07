package day7

import (
	"fmt"

	stl "aoc/stl"
)

func Solution() {
	ip := stl.ReadFile("input.txt")
	// ip := stl.ReadFile("example.txt")

	splits := countBeamSplits(ip)
	fmt.Printf("Tachyon beam split: %d times\n", splits)
}

func countBeamSplits(grid []string) (splits int) {
	nr, nc := len(grid), len(grid[0])

	visited := make([][]bool, 0, nr)
	for range nr {
		visited = append(visited, make([]bool, nc))
	}

	var dfs func(row, col int)
	dfs = func(row, col int) {
		if visited[row][col] {
			return
		}

		visited[row][col] = true
		if grid[row][col] == '^' {
			splits++
			if row+1 < nr {
				dfs(row+1, col-1)
				dfs(row+1, col+1)
			}
		} else {
			if row+1 < nr {
				dfs(row+1, col)
			}
		}
	}

	// 0, 7 is starting point in the example
	// 0, 70 is starting point in the example
	dfs(0, 70)

	return splits
}
