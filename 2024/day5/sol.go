package day5

import (
	"aoc/stl"
	"fmt"
)

func Solution() {
	ip := stl.ReadFile("input.txt")

	orders := [][]int{}
	for i, line := range ip {
		if len(line) == 0 {
			ip = ip[i+1:]
			break
		}
		orders = append(orders, stl.IntsFromString(ip[i]))
	}
	jobs := [][]int{}
	for _, job := range ip {
		jobs = append(jobs, stl.IntsFromString(job))
	}

	before := getorders(orders)
	inorderSum := inorder(jobs, before)
	fmt.Println(inorderSum)
}

func inorder(jobs [][]int, before map[int][]int) int {
	middleSum := 0
	for _, job := range jobs {
		outoforder := false
		for i, page := range job {
		each:
			for j := i + 1; j < len(job); j++ {
				for _, bef := range before[page] {
					if bef == job[j] {
						outoforder = true
						break each
					}
				}
			}
		}
		if !outoforder {
			middleSum += job[len(job)/2]
		}
	}
	return middleSum
}

func getorders(orders [][]int) map[int][]int {
	before := make(map[int][]int)
	for _, ord := range orders {
		first, second := ord[0], ord[1]
		before[second] = append(before[second], first)
	}
	return before
}
