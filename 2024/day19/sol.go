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
	for _, pattern := range patterns {
		curr := ""
		if form(pattern, curr, avail) {
			count++
		}
	}
	fmt.Println(count)
}

func form(target, curr string, avail []string) bool {
	if len(curr) > len(target) {
		return false
	}
	if curr == target {
		return true
	}
	for _, s := range avail {
		i, j := len(curr), len(s)
		if len(s) <= len(target[i:]) && target[i:i+j] == s {
			newS := curr + s
			if form(target, newS, avail) {
				return true
			}
		}
	}
	return false
}
