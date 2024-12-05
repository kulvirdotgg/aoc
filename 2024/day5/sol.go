package day5

import (
	"aoc/stl"
	"fmt"
)

func Solution() {
	ip := stl.ReadFile("input.txt")

	orders := [][]int{}
	jobs := [][]int{}
	isorder := true
	for i := range ip {
		if len(ip[i]) == 0 {
			isorder = false
			continue
		}
		if isorder {
			orders = append(orders, stl.IntsFromString(ip[i]))
		} else {
			jobs = append(jobs, stl.IntsFromString(ip[i]))
		}
	}

	orderMap := orderMap(orders)

	inorderSum := 0
	outoforderSum := 0
	for _, job := range jobs {
		if isinorder(job, orderMap) {
			inorderSum += job[len(job)/2]
		} else {
			outoforderSum += makeInorder(job, orderMap)
		}
	}
	fmt.Println(inorderSum)
	fmt.Println(outoforderSum)
}

func isinorder(job []int, orderMap map[int]map[int]bool) bool {
	for i, page := range job {
		for j := i + 1; j < len(job); j++ {
			if ok := orderMap[page][job[j]]; ok {
				return false
			}
		}
	}
	return true
}

func makeInorder(job []int, orderMap map[int]map[int]bool) int {
	ordered := make([]int, len(job)) // jobs in ordered sequence
	for i, num := range job {
		before := 0
		for j := 0; j < len(job); j++ {
			if j == i {
				continue
			}
			// count of elements that comes before ith elemnt
			// correct pos of i'th element will be num of elements [before] + 1
			if _, ok := orderMap[num][job[j]]; ok {
				before++
			}
		}
		ordered[before] = job[i]
		if ordered[len(job)/2] != 0 {
			return ordered[len(job)/2]
		}
	}
	return ordered[len(job)/2]
}

func orderMap(orders [][]int) map[int]map[int]bool {
	// 47|53 -> 47 comes before 53
	before := make(map[int]map[int]bool)
	for _, ord := range orders {
		first, second := ord[0], ord[1]
		if before[second] == nil {
			before[second] = make(map[int]bool)
		}
		// 53 : {{47: true}...}
		// elements coming before 53
		before[second][first] = true
	}
	return before
}
