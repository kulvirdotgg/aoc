package day13

import (
	"aoc/stl"
	"fmt"
)

const MAXINT int = 2147483647

func Solution() {
	ip := stl.ReadFile("input.txt")
	machines := getMachineMoves(ip)
	tokens := tokensForGifts(machines)
	fmt.Println(tokens)
}

func tokensForGifts(machines [][]pt) int {
	tokens := 0
	for _, machine := range machines {
		btn1, btn2, target := machine[0], machine[1], machine[2]
		least := MAXINT
		amul, bmul := 0, 0
		for amul*btn1.x <= target.x || amul*btn1.y <= target.y {
			bmul = (target.x - amul*btn1.x) / btn2.x

			// (A-move * btn1 + B-move * btn2) == target
			// A-move = A * btn1.x | A * btn1.y
			// B-move = B * btn2.x | B * btn2.y
			if amul*btn1.x+bmul*btn2.x == target.x && amul*btn1.y+bmul*btn2.y == target.y {
				cost := 3*amul + bmul
				least = min(least, cost)
			}
			amul++
		}
		if least != MAXINT {
			tokens += least
		}
	}
	return tokens
}

type pt struct {
	x, y int
}

func getMachineMoves(ip []string) [][]pt {
	machines := [][]pt{}
	idx := 0
	machine := []pt{}
	for _, line := range ip {
		if idx == 3 {
			machines = append(machines, machine)
			machine = []pt{}
			idx = 0
			continue
		}
		xy := stl.IntsFromString(line)
		pt := pt{x: xy[0], y: xy[1]}
		machine = append(machine, pt)
		idx++
	}
	machines = append(machines, machine)
	return machines
}
