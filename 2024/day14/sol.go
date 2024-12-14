package day14

import (
	"aoc/stl"
	"fmt"
)

func Solution() {
	ip := stl.ReadFile("input.txt")

	pv := [][]int{}
	for _, line := range ip {
		pv = append(pv, stl.IntsFromString(line))
	}

	safety := safetyFactor(pv)
	fmt.Println(safety)
}

const (
	W = 101
	H = 103
)

func easterEgg(pv [][]int) int {
	seconds := 0
	for {
		for _, robot := range pv {
			fmt.Println(robot)
		}
		seconds++
		break
	}
	return seconds
}

func safetyFactor(pv [][]int) int {
	quadCount := [4]int{}
	for _, robot := range pv {
		px, py := robot[0], robot[1]
		vx, vy := robot[2], robot[3]
		for range 100 {
			px = (px + vx + W) % W
			py = (py + vy + H) % H
		}
		if px < W/2 {
			if py < H/2 {
				quadCount[0]++
			} else if py > H/2 {
				quadCount[2]++
			}
		} else if px > W/2 {
			if py < H/2 {
				quadCount[1]++
			} else if py > H/2 {
				quadCount[3]++
			}
		}
	}

	safety := 1
	for _, cnt := range quadCount {
		safety *= cnt
	}
	return safety
}
