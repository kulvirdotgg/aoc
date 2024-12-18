package day18

import (
	"aoc/stl"
	"fmt"
)

const (
	N = 71
	M = 71
)

func Solution() {
	ip := stl.ReadFile("input.txt")

	matrix := [N][M]string{}
	for i := range N {
		for j := range M {
			matrix[i][j] = "."
		}
	}

	for _, line := range ip {
		pos := stl.IntsFromString(line)
		matrix[pos[0]][pos[1]] = "#"
		if bfs(matrix) == -1 {
			fmt.Println(pos)
			break
		}
	}
}

func bfs(matrix [N][M]string) int {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	queue := [][]int{}
	queue = append(queue, []int{0, 0})

	visited := [N][M]bool{}
	visited[0][0] = true
	level := 0
	for len(queue) != 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			top := queue[0]
			r, c := top[0], top[1]
			queue = queue[1:]

			if r == N-1 && c == M-1 {
				return level
			}

			for _, dir := range dirs {
				dx, dy := dir[0], dir[1]
				nr, nc := r+dx, c+dy

				if nr >= 0 && nr < N && nc >= 0 && nc < M && !visited[nr][nc] && matrix[nr][nc] != "#" {
					queue = append(queue, []int{nr, nc})
					visited[nr][nc] = true
				}
			}
		}
		level++
	}
	return -1
}

func constructMap(bytePos [][]int) [N][M]string {
	matrix := [N][M]string{}

	for i := range N {
		for j := range M {
			matrix[i][j] = "."
		}
	}

	for _, pos := range bytePos {
		x, y := pos[0], pos[1]
		matrix[x][y] = "#"
	}
	return matrix
}
