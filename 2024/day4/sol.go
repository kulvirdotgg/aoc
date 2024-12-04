package day4

import (
	"aoc/stl"
	"fmt"
)

func Solution() {
	ip := stl.ReadFile("input.txt")

	occurances := searchX(ip)
	fmt.Println(occurances)
}

func search(ws []string) int {
	rows, cols := len(ws), len(ws[0])
	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	word := "XMAS"
	getmax := func(r, c int, dir []int) bool {
		for dx := 1; dx < len(word); dx++ {
			nr := r + dir[0]*dx
			nc := c + dir[1]*dx

			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				return false
			}

			if ws[nr][nc] != word[dx] {
				return false
			}
		}
		return true
	}

	occurrences := 0
	for r := range rows {
		for c := range cols {
			letter := ws[r][c]
			if letter == 'X' {
				for _, dir := range directions {
					if getmax(r, c, dir) {
						occurrences++
					}
				}
			}
		}
	}
	return occurrences
}

func searchX(ws []string) int {
	rows, cols := len(ws), len(ws[0])

	d1 := [][]int{{-1, -1}, {1, 1}}
	d2 := [][]int{{-1, 1}, {1, -1}}
	getmax := func(r, c int) bool {
		d1x := ""
		d2x := ""
		for _, d := range d1 {
			nr := r + d[0]
			nc := c + d[1]

			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				return false
			}
			d1x += string(ws[nr][nc])
		}

		for _, d := range d2 {
			nr := r + d[0]
			nc := c + d[1]

			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				return false
			}
			d2x += string(ws[nr][nc])
		}
		if (d1x == "MS" || d1x == "SM") && (d2x == "MS" || d2x == "SM") {
			return true
		}
		return false
	}

	occurrences := 0
	for r := range rows {
		for c := range cols {
			if ws[r][c] == 'A' {
				if getmax(r, c) {
					occurrences++
				}
			}
		}
	}
	return occurrences
}
