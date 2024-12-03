package day3

import (
	stl "aoc/stl"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Solution() {
	ip := stl.ReadFile("input.txt")

	sop := 0
	for _, line := range ip {
		sop += mull(line)
	}
	fmt.Println(sop)

	// Have to do this because, at the begining, `do` is true
	// but that will change for next lines
	// Next line state will be effected by previous lines
	// Hence if there is don't() in previous line, next line will begin dont()
	instructions := [][]string{}
	for _, line := range ip {
		instructions = append(instructions, extractInstruction(line)...)
	}
	sopdodont := mullDoDont(instructions)
	fmt.Println(sopdodont)
}

func mull(line string) int {
	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)
	instructions := r.FindAllStringSubmatch(line, -1)

	sam := 0 // sumAfterMultiplications
	for _, instruction := range instructions {
		num1, _ := strconv.Atoi(instruction[1])
		num2, _ := strconv.Atoi(instruction[2])

		sam += num1 * num2
	}
	return sam
}

func extractInstruction(line string) [][]string {
	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)|don\'t\(\)|do\(\)`)
	return r.FindAllStringSubmatch(line, -1)
}

func mullDoDont(instructions [][]string) int {
	sam := 0
	doMul := true
	for _, instruction := range instructions {
		if strings.Contains(instruction[0], "don't()") {
			doMul = false
			continue
		} else if strings.Contains(instruction[0], "do()") {
			doMul = true
			continue
		}

		if doMul {
			num1, _ := strconv.Atoi(instruction[1])
			num2, _ := strconv.Atoi(instruction[2])
			sam += num1 * num2
		}
	}
	return sam
}
