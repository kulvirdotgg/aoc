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

	ways = countWaysWithDAC(adj)
	fmt.Printf("total number of way from 'srv' to 'out' with 'dac and 'fft': %d\n", ways)
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

func countWaysWithDAC(adj map[string][]string) (ways int) {
	visited := make(map[string]int)

	var dfs func(node string, dac bool, fft bool) int
	dfs = func(node string, dac bool, fft bool) int {
		key := fmt.Sprintf("%s:%v%v", node, dac, fft)

		if cnt, ok := visited[key]; ok {
			return cnt
		}

		if node == "out" {
			// it will be valid path only if both "fft" and "dac" nodes
			if dac && fft {
				visited[key] = 1
				return 1
			}
			visited[key] = 0
			return 0
		}

		for _, conn := range adj[node] {
			visited[key] += dfs(conn, dac || conn == "dac", fft || conn == "fft")
		}
		return visited[key]
	}

	ways = dfs("svr", false, false)
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
