package day19

import (
	"aoc/stl"
	"fmt"
	"sort"
	"strings"
)

func Solution() {
	ip := stl.ReadFile("input.txt")

	avail := strings.Split(ip[0], ", ")
	sort.Slice(avail, func(i, j int) bool {
		return avail[i] < avail[j]
	})
	patterns := ip[2:]

	count := 0
	totalCount := 0
	dp := make(map[string]int)
	for _, pattern := range patterns {
		if form(pattern, 0, avail) {
			count++
		}
		totalCount += formWithCount(pattern, avail, dp)
	}
	fmt.Println(count)
	fmt.Println(totalCount)
}

func form(target string, idx int, avail []string) bool {
	if idx == len(target) {
		return true
	}

	for _, s := range avail {
		if strings.HasPrefix(target[idx:], s) {
			if form(target, idx+len(s), avail) {
				return true
			}
		}
	}
	return false
}

func formWithCount(target string, avail []string, dp map[string]int) int {
	if val, ok := dp[target]; ok {
		return val
	}

	total := 0
	for _, s := range avail {
		if strings.Index(target, s) == 0 {
			if len(s) == len(target) {
				total++
				continue
			}
			total += formWithCount(target[len(s):], avail, dp)
		}
	}
	dp[target] = total
	return total
}
