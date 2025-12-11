package day10

import (
	"fmt"
	"strings"

	"aoc/stl"
)

// https://adventofcode.com/2025/day/11
func Solution() {
	// ip := stl.ReadFile("example.txt")
	ip := stl.ReadFile("input.txt")
	adj := adjMatrix(ip)

	ways := countWays(adj)
	fmt.Printf("total number of way from 'you' to 'out': %d\n", ways)
}

func countWays(adj map[string][]string) (ways int) {
	visited := make(map[string]int)

	var dfs func(node string) int
	dfs = func(node string) int {
		if visited[node] > 0 {
			return visited[node]
		}

		if node == "out" {
			return 1
		}

		ways := 0
		for _, conn := range adj[node] {
			ways += dfs(conn)
		}
		visited[node] = ways
		return ways
	}

	ways = dfs("you")
	return
}

func adjMatrix(puzzle []string) map[string][]string {
	adj := make(map[string][]string)

	for _, row := range puzzle {
		devices := strings.Split(row, " ")
		// "aaa:" -> "aaa"
		devices[0] = devices[0][:len(devices[0])-1]

		adj[devices[0]] = append(adj[devices[0]], devices[1:]...)
	}

	return adj
}
