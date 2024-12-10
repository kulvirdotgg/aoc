package day10

import (
	"aoc/stl"
	"fmt"
)

func Solution() {
	ip := stl.ReadFile("input.txt")

	topoMap := stl.MatrixFromString(ip)
	trails := 0
	for r := range topoMap {
		for c := range topoMap {
			if topoMap[r][c] == 0 {
				trails += bfs(topoMap, r, c)
			}
		}
	}
	fmt.Println(trails)
}

type point struct {
	x, y int
}

func bfs(topoMap [][]int64, r, c int) int {
	rows, cols := len(topoMap), len(topoMap[0])

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	uniqEnds := make(map[point]struct{})

	q := [][]int{}
	pt := []int{r, c}
	q = append(q, pt)

	for len(q) != 0 {
		pos := q[0]
		q = q[1:]

		if topoMap[pos[0]][pos[1]] == 9 {
			// fmt.Println(topoMap[pos[0]][pos[1]], pos)
			uniqEnds[point{x: pos[0], y: pos[1]}] = struct{}{}
			continue
		}

		for _, dir := range dirs {
			nr, nc := pos[0]+dir[0], pos[1]+dir[1]

			if nr >= rows || nr < 0 || nc >= cols || nc < 0 {
				continue
			}

			if topoMap[nr][nc] == topoMap[pos[0]][pos[1]]+1 {
				pt := []int{nr, nc}
				q = append(q, pt)
			}
		}
	}
	return len(uniqEnds)
}
