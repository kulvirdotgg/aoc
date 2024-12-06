package day6

import (
	stl "aoc/stl"
	"fmt"
)

func Solution() {
	ip := stl.ReadFile("extc.txt")

	visited := countMoves(ip)
	fmt.Println(visited)
}

func countMoves(ip []string) int {
	rows, cols := len(ip), len(ip[0])
	row, col := -1, -1 // the position of the guard
	for r := range ip {
		for c := range ip[r] {
			if ip[r][c] == '^' {
				row, col = r, c
			}
		}
	}

	// fuck you golang
	visited := make([][]bool, len(ip))
	for i := range visited {
		visited[i] = make([]bool, len(ip[i]))
	}

	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	d := 0
	r, c := row, col
	for r < rows && r >= 0 && c < cols && c >= 0 {
		visited[r][c] = true
		nr := r + directions[d][0]
		nc := c + directions[d][1]

		if nr < rows && nr >= 0 && nc < cols && nc >= 0 && ip[nr][nc] == '#' {
			d = (d + 1) % 4
			nr, nc = r+directions[d][0], c+directions[d][1]
		}
		r, c = nr, nc
	}

	visCnt := 0
	for r := range visited {
		for c := range visited[r] {
			if visited[r][c] == true {
				visCnt++
			}
		}
	}
	return visCnt
}
