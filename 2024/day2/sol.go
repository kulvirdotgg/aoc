package day2

import (
	stl "aoc/stl"
	"fmt"
	"slices"
)

func Solution() {
	ip := stl.ReadFile("input.txt")
	safeCount := 0
	for _, report := range ip {
		ints := stl.IntsFromString(report)
		if areYouSure(ints) {
			safeCount++
		}
	}
	fmt.Println(safeCount)
}

func safe(ints []int) bool {
	incr := true // by default assume sequence to be increasing
	if ints[0] > ints[1] {
		incr = false
	}

	for i := 1; i < len(ints); i++ {
		diff := ints[i] - ints[i-1]
		if diff < 0 {
			diff *= -1
		}

		if (incr && ints[i-1] > ints[i]) || (!incr && ints[i-1] < ints[i]) || diff > 3 || diff == 0 {
			return false
		}
	}
	return true
}

func areYouSure(ints []int) bool {
	if safe(ints) {
		return true
	}

	for lidx := 0; lidx < len(ints); lidx++ {
		slice := []int{}
		slice = append(slice, ints...)
		slice = slices.Delete(slice, lidx, lidx+1)

		if safe(slice) {
			return true
		}
	}
	return false
}
