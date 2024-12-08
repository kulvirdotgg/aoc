package day8

import (
	"aoc/stl"
	"fmt"
)

func Solution() {
	ip := stl.ReadFile("input.txt")
	antinodes := antiNodes(ip)
	fmt.Println(antinodes)
}

type coord struct {
	x, y int
}

func antiNodes(ip []string) int {
	freq := make(map[rune][]coord)
	for x, line := range ip {
		for y, ch := range line {
			if ch != '.' {
				if _, ok := freq[ch]; !ok {
					freq[ch] = make([]coord, 0)
				}
				c := coord{x, y}
				freq[ch] = append(freq[ch], c)
			}
		}
	}

	distance := func(c1, c2 coord) (int, int) {
		return c2.x - c1.x, c2.y - c1.y
	}

	validNode := func(antinode coord) bool {
		return antinode.x < len(ip) && antinode.x >= 0 && antinode.y < len(ip[0]) && antinode.y >= 0
	}

	antinodes := make(map[coord]struct{})
	for _, antennas := range freq {
		for i := 0; i < len(antennas)-1; i++ {
			for j := i + 1; j < len(antennas); j++ {
				dx, dy := distance(antennas[i], antennas[j])

				anti1 := coord{x: antennas[i].x - dx, y: antennas[i].y - dy}
				if validNode(anti1) {
					antinodes[anti1] = struct{}{}
				}

				anti2 := coord{x: antennas[j].x + dx, y: antennas[j].y + dy}
				if validNode(anti2) {
					antinodes[anti2] = struct{}{}
				}
			}
		}
	}
	return len(antinodes)
}
