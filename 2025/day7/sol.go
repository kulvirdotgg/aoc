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

	beams := countBeams(ip)
	fmt.Printf("Total paths taken by Tachyon beam: %d times\n", beams)
}

func countBeamSplits(grid []string) (splits int) {
	nr, nc := len(grid), len(grid[0])

	visited := make([][]bool, 0, nr)
	for range nr {
		visited = append(visited, make([]bool, nc))
	}

	var dfs func(row, col int)
	dfs = func(row, col int) {
		if nr <= row {
			return
		}

		if visited[row][col] {
			return
		}

		visited[row][col] = true
		if grid[row][col] == '^' {
			splits++
			dfs(row+1, col-1)
			dfs(row+1, col+1)
		} else {
			dfs(row+1, col)
		}
	}

	dfs(0, 70)

	return splits
}

func countBeams(grid []string) (beams int) {
	nr, nc := len(grid), len(grid[0])

	visited := make([][]int, 0, nr)
	for range nr {
		visited = append(visited, make([]int, nc))
	}

	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if nr <= row {
			return 1
		}

		if visited[row][col] != 0 {
			return visited[row][col]
		}

		if grid[row][col] == '^' {
			visited[row][col] = dfs(row+1, col-1) + dfs(row+1, col+1)
		} else {
			visited[row][col] = dfs(row+1, col)
		}
		return visited[row][col]
	}

	beams = dfs(0, 70)

	return beams
}
