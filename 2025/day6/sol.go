package day6

import (
	"fmt"

	stl "aoc/stl"
)

// https://adventofcode.com/2025/day/6
func Solution() {
	ip := stl.ReadFile("input.txt")
	// ip := stl.ReadFile("example.txt")
	nums, ops := getNumsAndOps(ip)
	total := solveProblems(nums, ops)
	fmt.Println(total)
}

func solveProblems(nums [][]int, ops []string) (total int) {
	for col, op := range ops {
		switch op {
		case "+":
			acc := 0
			for row := range len(nums) {
				acc += nums[row][col]
			}
			total += acc
		case "*":
			acc := 1
			for row := range len(nums) {
				acc *= nums[row][col]
			}
			total += acc
		}
	}
	return total
}

func getNumsAndOps(ip []string) (nums [][]int, ops []string) {
	n := len(ip)

	opsRow := ip[n-1]
	for _, ch := range opsRow {
		if ch != ' ' {
			ops = append(ops, string(ch))
		}
	}

	numsRows := ip[:n-1]
	nums = make([][]int, 0, len(numsRows))
	for _, row := range numsRows {
		ints := stl.IntsFromString(row)
		nums = append(nums, ints)
	}

	return nums, ops
}
