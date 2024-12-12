package day12

import (
	"aoc/stl"
	"fmt"
)

func Solution() {
	ip := stl.ReadFile("input.txt")
	price := fencingPrice(ip)
	fmt.Println(price)
}

func fencingPrice(ip []string) int {
	visited := make([][]bool, len(ip))
	for r := range visited {
		visited[r] = make([]bool, len(ip[0]))
	}

	price := 0
	for i, line := range ip {
		for j := range line {
			if visited[i][j] == false {
				area, peri := dfs(ip, i, j, visited)
				price += area * peri
			}
		}
	}
	return price
}

var dirs [][]int = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func dfs(ip []string, r, c int, visited [][]bool) (int, int) {
	area, peri := 1, 0
	visited[r][c] = true

	for _, dir := range dirs {
		dx, dy := dir[0], dir[1]
		nr, nc := r+dx, c+dy
		if nr < len(ip) && nr >= 0 && nc < len(ip[0]) && nc >= 0 && ip[nr][nc] == ip[r][c] {
			if !visited[nr][nc] {
				na, np := dfs(ip, nr, nc, visited)
				area += na
				peri += np
			}
		} else {
			peri += 1
		}
	}
	return area, peri
}
