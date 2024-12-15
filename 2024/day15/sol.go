package day15

import (
	"aoc/stl"
	"fmt"
)

func Solution() {
	ip := stl.ReadFile("extc.txt")

	moves := []string{}
	naksha := []string{}
	for i, line := range ip {
		if len(line) == 0 {
			naksha = ip[:i]
			moves = ip[i+1:]
			break
		}
	}
	newMap := goRobo(moves, getMap(naksha))
	for _, line := range newMap {
		fmt.Println(line)
	}
	distance := lanternGps(newMap)
	fmt.Println(distance)
}

func lanternGps(naksha [][]string) int {
	distance := 0
	for r, line := range naksha {
		for c, ch := range line {
			if ch == "O" {
				distance += 100*r + c
			}
		}
	}
	return distance
}

func goRobo(moves []string, naksha [][]string) [][]string {
	locateBot := func() [2]int {
		robo := [2]int{}
		for r, line := range naksha {
			for c, ch := range line {
				if ch == "@" {
					robo[0] = r
					robo[1] = c
					return robo
				}
			}
		}
		return robo
	}

	for _, str := range moves {
		for _, mv := range str {
			robo := locateBot()
			switch {
			case mv == '<':
				naksha = move(naksha, 0, robo[0], robo[1])
			case mv == '^':
				naksha = move(naksha, 1, robo[0], robo[1])
			case mv == '>':
				naksha = move(naksha, 2, robo[0], robo[1])
			case mv == 'v':
				naksha = move(naksha, 3, robo[0], robo[1])
			}
		}
	}
	return naksha
}

func move(naksha [][]string, dir, row, col int) [][]string {
	dirs := [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

	nr, nc := row+dirs[dir][0], col+dirs[dir][1]
	if naksha[nr][nc] == "#" {
		return naksha
	}

	box := false
	if naksha[nr][nc] == "O" {
		box = true
	}

	nnr, nnc := nr, nc
	if box {
		for naksha[nnr][nnc] == "O" {
			nnr += dirs[dir][0]
			nnc += dirs[dir][1]
		}
		if naksha[nnr][nnc] != "#" {
			naksha[nnr][nnc] = "O"
			naksha[nr][nc] = "@"
			naksha[row][col] = "."
		}
	} else {
		naksha[nr][nc] = "@"
		naksha[row][col] = "."
	}
	return naksha
}

func getMap(naksha []string) [][]string {
	newMap := make([][]string, len(naksha))
	for r, line := range naksha {
		newMap[r] = make([]string, 0)
		for _, ch := range line {
			newMap[r] = append(newMap[r], string(ch))
		}
	}
	return newMap
}
